package dsg

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/log"
	"math/big"
	"sync"
)

type ValidationResult int

const (
	inmMemoryProposals  = 96   // Number of recent block proposals to keep in memory
	inMemoryValidations = 4096 // Number of recent validation messages to keep in memory
	inMemoryReqNewBPs   = 4096 // Number of recent request new block proposal messages to keep in memory

	Accept ValidationResult = 1
	Reject ValidationResult = -1
	Unknown ValidationResult = 0
)

type Cache struct {
	validations     *lru.ARCCache
	proposals       *lru.ARCCache
	reqNewBPs       *lru.ARCCache

	invalidCounter   uint64 // counter for tracking invalid BPs, including BP/VM msg timeouts. Resets with each block commit
	invalidCounterMu sync.RWMutex
}

type ValidationKey struct {
	BlockNum  *big.Int         `json:"blocknum"`
	Proposer  common.Address   `json:"proposer"`
	Validator common.Address   `json:"validator"`
}

type ProposalKey struct {
	BlockNum uint64     `json:"blocknum"`
	Proposer common.Address  `json:"proposer"`
}

type ReqNewBPKey struct {
	BlockNum  *big.Int       `json:"blocknum"`
	Validator common.Address `json:"validator"`
}

func NewCache() *Cache {

	validations, _ := lru.NewARC(inMemoryValidations)
	proposals, _ := lru.NewARC(inmMemoryProposals)
	reqNewBPs, _ := lru.NewARC(inMemoryReqNewBPs)

	cache := &Cache{
		validations:    validations,
		proposals:      proposals,
		reqNewBPs:      reqNewBPs,
		invalidCounter: 0,
	}
	return cache
}

func (c *Cache) InsertBlockProposal(bp BlockProposal) {
	key := ProposalKey{
		BlockNum: bp.Number.Uint64(),
		Proposer: bp.Proposer,
	}

	if ok := c.proposals.Contains(key); !ok {
		log.Info("cache new block proposal", "num", bp.Number, "proposer", bp.Proposer)
		c.proposals.Add(key, bp)
	}
}

func (c *Cache) GetBlockProposal(blockNum *big.Int, proposer common.Address) (BlockProposal, error) {
	key := ProposalKey{
		BlockNum: blockNum.Uint64(),
		Proposer: proposer,
	}

	var blockProposal BlockProposal
	if bp, ok := c.proposals.Get(key); ok {
		if blockProposal, ok = bp.(BlockProposal); ok {
			return blockProposal, nil
		}
	}

	return blockProposal, errGetBPFromCacheFailed
}

func (c *Cache) GetBlockProposalByHash(hash common.Hash) (BlockProposal, error) {
	var blockProposal BlockProposal
	for _, key := range c.proposals.Keys() {
		if bp, ok := c.proposals.Get(key); ok {
			if blockProposal, ok = bp.(BlockProposal); ok {
				if blockProposal.BlockHash == hash {
					return blockProposal, nil
				}
			}
		}
	}

	return blockProposal, errGetBPFromCacheFailed
}

func (c *Cache) InsertValidationMessage(msg ValidationMessage) ValidationResult {
	return c.insertValidationMessage(msg)
}

func (c *Cache) insertValidationMessage(msg ValidationMessage) ValidationResult {
	n := msg.Number
	v := msg.Verifier
	p := msg.Proposer

	key := ValidationKey{
		BlockNum:  n,
		Validator: v,
		Proposer:  p,
	}

	if ok := c.validations.Contains(key); ok {
		log.Info("delete stale validation message from cache", "num", msg.Number, "proposer", msg.Proposer, "validator", msg.Verifier)
		c.validations.Remove(key)
	}

	log.Info("cache new validation message", "num", msg.Number, "proposer", msg.Proposer, "validator", msg.Verifier)
	c.validations.Add(key, msg)

	return c.QueryValidationState(n, p)
}

func (c *Cache) QueryValidationState(blockNum *big.Int, proposer common.Address) ValidationResult {
	acks := float64(0)
	nacks := float64(0)

	var validationMessage ValidationMessage
	for _, key := range c.validations.Keys() {
		if vm, ok := c.validations.Get(key); ok {
			if validationMessage, ok = vm.(ValidationMessage); ok {
				if (validationMessage.Number.Cmp(blockNum) == 0) && (validationMessage.Proposer == proposer) && (validationMessage.Authorize == true) {
					acks = acks + 1
				}
				if (validationMessage.Number.Cmp(blockNum) == 0) && (validationMessage.Proposer == proposer) && (validationMessage.Authorize == false) {
					nacks = nacks + 1
				}
			}
		}
	}

	ackRequirement, nackRequirement, acknowledges, rejections := c.calculateAcksNacks(common.NumSignersInRound, acks, nacks)

	if acknowledges >= ackRequirement {
		return Accept
	}
	if rejections >= nackRequirement {
		return Reject
	}
	return Unknown
}

// CalculateAcksNacks is the exported function for calculateAcksNacks
func (c *Cache) CalculateAcksNacks(numSigners int, acks, nacks float64) (float64, float64, float64, float64) {
	return c.calculateAcksNacks(numSigners, acks, acks)
}

// calculateAcksNacks calculates and returns the ACK/NACK requirements and
// the actual ACK/NACK values
func (c *Cache) calculateAcksNacks(numSigners int, acks, nacks float64) (float64, float64, float64, float64) {

	acknowledges, rejections := float64(0), float64(0)
	totalSignersFloat := float64(numSigners)

	ackRequirement, nackRequirement := F()

	if acks > 0 {
		acknowledges = acks / totalSignersFloat
	}
	if nacks > 0 {
		rejections = nacks / totalSignersFloat
	}

	return ackRequirement, nackRequirement, acknowledges, rejections
}

// IncrementInvalidCounter is used to increase the timeoutCounter for a block.
// It can be called when an EV times out waiting for a new BP, when a BP is
// considered invalid according to consensus, or when waiting for 2/3 ACKs times out
func (c *Cache) IncrementInvalidCounter() {
	c.invalidCounterMu.Lock()
	defer c.invalidCounterMu.Unlock()

	c.invalidCounter = c.invalidCounter + 1
}

// ResetInvalidCounter is used to reset the timeoutCounter when a block has been committed
// so the tracking can begin again for the next block proposal(s)
func (c *Cache) ResetInvalidCounter() {
	c.invalidCounterMu.Lock()
	defer c.invalidCounterMu.Unlock()

	c.invalidCounter = 0
}

// GetInvalidCounter is used to get the current timeoutCounter value
func (c *Cache) GetInvalidCounter() uint64 {
	c.invalidCounterMu.RLock()
	defer c.invalidCounterMu.RUnlock()

	return c.invalidCounter
}

// InsertRequestNewBlockProposalMessage inserts a Request New Block Proposal message
// into the cache, and executes QueryRequestNewBpMessageState to see if
// 2/3 EVs agree a new BP is required
func (c *Cache) InsertRequestNewBlockProposalMessage(msg RequestNewBlockProposalMessage) bool {
	c.insertRequestNewBlockProposalMessage(msg)
	return c.QueryRequestNewBpMessageState(msg.Number)
}

// QueryRequestNewBpMessageState queries the reqNewBPs cache for the given
// block. Returns true when 2/3 EVs have sent a new request
func (c *Cache) QueryRequestNewBpMessageState(blockNum *big.Int) bool {
	numRequests := float64(0)

	var reqNewBPMessage RequestNewBlockProposalMessage
	for _, key := range c.reqNewBPs.Keys() {
		if reqm, ok := c.reqNewBPs.Get(key); ok {
			if reqNewBPMessage, ok = reqm.(RequestNewBlockProposalMessage); ok {
				if reqNewBPMessage.Number.Cmp(blockNum) == 0{
					numRequests = numRequests + 1
				}
			}
		}
	}

	requirement, _, requests, _ := c.calculateAcksNacks(common.NumSignersInRound, numRequests, float64(0))

	if requests >= requirement {
		return true
	}

	return false
}

func (c *Cache) insertRequestNewBlockProposalMessage(msg RequestNewBlockProposalMessage) {
	key := ReqNewBPKey{
		BlockNum:  msg.Number,
		Validator: msg.Verifier,
	}
	if ok := c.reqNewBPs.Contains(key); ok {
		log.Info("delete stale request new block proposal message from cache", "num", msg.Number, "validator", msg.Verifier)
		c.reqNewBPs.Remove(key)
	}
	log.Info("cache new request new block proposal message", "num", msg.Number, "validator", msg.Verifier)
	c.reqNewBPs.Add(key, msg)
}
