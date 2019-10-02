package dsg

import (
	"errors"
	lru "github.com/hashicorp/golang-lru"
	"github.com/unification-com/mainchain/common"
	"math/big"
	"sync"
)

type ValidationResult int

const (
	inmMemoryProposals  = 96   // Number of recent block proposals to keep in memory
	inMemoryValidations = 4096 // Number of recent validation messages to keep in memory

	Accept ValidationResult = 1
	Reject ValidationResult = -1
	Unknown ValidationResult = 0
)

type Cache struct {
	validations     *lru.ARCCache
	proposals       *lru.ARCCache

	invalidCounter   uint64 // counter for tracking invalid BPs, including BP/VM msg timeouts. Resets with each block commit
	invalidCounterMu sync.RWMutex
}

type ValidationKey struct {
	BlockNum  uint64         `json:"blocknum"`
	Proposer  common.Address `json:"proposer"`
	Validator common.Address `json:"validator"`
}

type ProposalKey struct {
	BlockNum uint64         `json:"blocknum"`
	Proposer common.Address `json:"proposer"`
}

func NewCache() *Cache {

	validations, _ := lru.NewARC(inMemoryValidations)
	proposals, _ := lru.NewARC(inmMemoryProposals)

	cache := &Cache{
		validations:    validations,
		proposals:      proposals,
		invalidCounter: 0,
	}
	return cache
}

func (d *Cache) InsertBlockProposal(bp BlockProposal) {
	n := bp.Number.Uint64()
	p := bp.Proposer

	key := ProposalKey{
		BlockNum: n,
		Proposer: p,
	}

	d.proposals.Add(key, bp)
}

func (d *Cache) GetBlockProposal(blockNum *big.Int, proposer common.Address) (BlockProposal, error) {
	key := ProposalKey{
		BlockNum: blockNum.Uint64(),
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
	n := msg.Number.Uint64()
	v := msg.Verifier
	p := msg.Proposer

	key := ValidationKey{
		BlockNum:  n,
		Validator: v,
		Proposer:  p,
	}

	d.validations.Add(key, msg)

	return d.acceptBlock(n, p)
}

func (d *Cache) acceptBlock(blockNum uint64, proposer common.Address) ValidationResult {
	acks := float64(0)
	nacks := float64(0)

	var validationMessage ValidationMessage
	for _, key := range d.validations.Keys() {
		if vm, ok := d.validations.Get(key); ok {
			if validationMessage, ok = vm.(ValidationMessage); ok {
				if (validationMessage.Number.Uint64() == blockNum) && (validationMessage.Proposer == proposer) && (validationMessage.Authorize == true) {
					acks = acks + 1
				}
				if (validationMessage.Number.Uint64() == blockNum) && (validationMessage.Proposer == proposer) && (validationMessage.Authorize == false) {
					nacks = nacks + 1
				}
			}
		}
	}

	totalSignersFloat := float64(common.NumSignersInRound)
	ackRequirement := float64(2) / totalSignersFloat
	nackRequirement := float64(1) / totalSignersFloat
	acknowledges := acks / totalSignersFloat
	rejections := nacks / totalSignersFloat

	if acknowledges >= ackRequirement {
		return Accept
	}
	if rejections >= nackRequirement {
		return Reject
	}
	return Unknown
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
