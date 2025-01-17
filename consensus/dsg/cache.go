package dsg

import (
	"errors"
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
	BlockNum common.Hash     `json:"blocknum"`
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

func (d *Cache) InsertBlockProposal(bp BlockProposal) {
	key := ProposalKey{
		BlockNum: common.BytesToHash(bp.Number.Bytes()),
		Proposer: bp.Proposer,
	}

	if ok := d.proposals.Contains(key); !ok {
		log.Info("cache new block proposal", "num", bp.Number, "proposer", bp.Proposer)
		d.proposals.Add(key, bp)
	}
}

func (d *Cache) GetBlockProposal(blockNum *big.Int, proposer common.Address) (BlockProposal, error) {
	key := ProposalKey{
		BlockNum: common.BytesToHash(blockNum.Bytes()),
		Proposer: proposer,
	}

	var blockProposal BlockProposal
	if bp, ok := d.proposals.Get(key); ok {
		if blockProposal, ok = bp.(BlockProposal); ok {
			return blockProposal, nil
		}
	}

	return blockProposal, errors.New("could not retrieve block proposal from cache")
}

func (d *Cache) GetBlockProposalByHash(hash common.Hash) (BlockProposal, error) {
	var blockProposal BlockProposal
	for _, key := range d.proposals.Keys() {
		if bp, ok := d.proposals.Get(key); ok {
			if blockProposal, ok = bp.(BlockProposal); ok {
				if blockProposal.BlockHash == hash {
					return blockProposal, nil
				}
			}
		}
	}

	return blockProposal, errors.New("could not retrieve block proposal from cache")
}

func (d *Cache) InsertValidationMessage(msg ValidationMessage) ValidationResult {
	return d.insertValidationMessage(msg)
}

func (d *Cache) insertValidationMessage(msg ValidationMessage) ValidationResult {
	n := msg.Number
	v := msg.Verifier
	p := msg.Proposer

	key := ValidationKey{
		BlockNum:  n,
		Validator: v,
		Proposer:  p,
	}

	if ok := d.validations.Contains(key); ok {
		log.Info("delete stale validation message from cache", "num", msg.Number, "proposer", msg.Proposer, "validator", msg.Verifier)
		d.validations.Remove(key)
	}

	log.Info("cache new validation message", "num", msg.Number, "proposer", msg.Proposer, "validator", msg.Verifier)
	d.validations.Add(key, msg)

	return d.QueryValidationState(n, p)
}

func (d *Cache) QueryValidationState(blockNum *big.Int, proposer common.Address) ValidationResult {
	acks := float64(0)
	nacks := float64(0)

	var validationMessage ValidationMessage
	for _, key := range d.validations.Keys() {
		if vm, ok := d.validations.Get(key); ok {
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

	ackRequirement, nackRequirement, acknowledges, rejections := d.calculateAcksNacks(common.NumSignersInRound, acks, nacks)

	if acknowledges >= ackRequirement {
		return Accept
	}
	if rejections >= nackRequirement {
		return Reject
	}
	return Unknown
}

// CalculateAcksNacks is the exported function for calculateAcksNacks
func (d *Cache) CalculateAcksNacks(numSigners int, acks, nacks float64) (float64, float64, float64, float64) {
	return d.calculateAcksNacks(numSigners, acks, acks)
}

// calculateAcksNacks calculates and returns the ACK/NACK requirements and
// the actual ACK/NACK values
func (d *Cache) calculateAcksNacks(numSigners int, acks, nacks float64) (float64, float64, float64, float64) {

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
func (d *Cache) IncrementInvalidCounter() {
	d.invalidCounterMu.Lock()
	defer d.invalidCounterMu.Unlock()

	d.invalidCounter = d.invalidCounter + 1
}

// ResetInvalidCounter is used to reset the timeoutCounter when a block has been committed
// so the tracking can begin again for the next block proposal(s)
func (d *Cache) ResetInvalidCounter() {
	d.invalidCounterMu.Lock()
	defer d.invalidCounterMu.Unlock()

	d.invalidCounter = 0
}

// GetInvalidCounter is used to get the current timeoutCounter value
func (d *Cache) GetInvalidCounter() uint64 {
	d.invalidCounterMu.RLock()
	defer d.invalidCounterMu.RUnlock()

	return d.invalidCounter
}

// InsertRequestNewBlockProposalMessage inserts a Request New Block Proposal message
// into the cache, and executes QueryRequestNewBpMessageState to see if
// 2/3 EVs agree a new BP is required
func (d *Cache) InsertRequestNewBlockProposalMessage(msg RequestNewBlockProposalMessage) bool {
	d.insertRequestNewBlockProposalMessage(msg)
	return d.QueryRequestNewBpMessageState(msg.Number)
}

// QueryRequestNewBpMessageState queries the reqNewBPs cache for the given
// block. Returns true when 2/3 EVs have sent a new request
func (d *Cache) QueryRequestNewBpMessageState(blockNum *big.Int) bool {
	numRequests := float64(0)

	var reqNewBPMessage RequestNewBlockProposalMessage
	for _, key := range d.reqNewBPs.Keys() {
		if reqm, ok := d.reqNewBPs.Get(key); ok {
			if reqNewBPMessage, ok = reqm.(RequestNewBlockProposalMessage); ok {
				if reqNewBPMessage.Number.Cmp(blockNum) == 0{
					numRequests = numRequests + 1
				}
			}
		}
	}

	requirement, _, requests, _ := d.calculateAcksNacks(common.NumSignersInRound, numRequests, float64(0))

	if requests >= requirement {
		return true
	}

	return false
}

func (d *Cache) insertRequestNewBlockProposalMessage(msg RequestNewBlockProposalMessage) {
	key := ReqNewBPKey{
		BlockNum:  msg.Number,
		Validator: msg.Verifier,
	}
	if ok := d.reqNewBPs.Contains(key); ok {
		log.Info("delete stale request new block proposal message from cache", "num", msg.Number, "validator", msg.Verifier)
		d.reqNewBPs.Remove(key)
	}
	log.Info("cache new request new block proposal message", "num", msg.Number, "validator", msg.Verifier)
	d.reqNewBPs.Add(key, msg)
}
