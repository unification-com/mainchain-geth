package dsg

import (
	"errors"
	lru "github.com/hashicorp/golang-lru"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/log"
	"math/big"
	"time"
)

const (
	inmMemoryProposals  = 96   // Number of recent block proposals to keep in memory
	inMemoryValidations = 4096 // Number of recent validation messages to keep in memory
)

type Cache struct {
	validations *lru.ARCCache
	proposals   *lru.ARCCache
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
		validations: validations,
		proposals:   proposals,
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

	log.Info("cache new block proposal", "block", key.BlockNum, "proposer", key.Proposer)

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

func (d *Cache) InsertValidationMessage(msg ValidationMessage) bool {
	return d.insertValidationMessage(msg)
}

func (d *Cache) insertValidationMessage(msg ValidationMessage) bool {
	n := msg.Number.Uint64()
	v := msg.Verifier
	p := msg.Proposer

	key := ValidationKey{
		BlockNum:  n,
		Validator: v,
		Proposer:  p,
	}

	log.Info("insertValidationMessage", "block", key.BlockNum, "proposer", key.Proposer, "validator", key.Validator, "authorise", msg.Authorize)
	d.validations.Add(key, msg)

	return d.acceptBlock(n, p)
}

func (d *Cache) acceptBlock(blockNum uint64, proposer common.Address) bool {
	acks := float64(0)

	var validationMessage ValidationMessage
	for _, key := range d.validations.Keys() {
		if vm, ok := d.validations.Get(key); ok {
			if validationMessage, ok = vm.(ValidationMessage); ok {
				if (validationMessage.Number.Uint64() == blockNum) && (validationMessage.Proposer == proposer) && (validationMessage.Authorize == true) {
					acks = acks + 1
				}
			}
		}
	}

	totalSignersFloat := float64(common.NumSignersInRound)
	requirement := float64(2) / totalSignersFloat
	acknowledges := acks / totalSignersFloat
	log.Info("acceptBlock", "num", blockNum, "proposer", proposer, "acks", acks)
	return acknowledges >= requirement
}

func (d *Cache) PollBlockProposalCache(blockNum *big.Int, proposer common.Address) (BlockProposal, error) {
	var blockProposal BlockProposal
	log.Info("PollBlockProposalCache", "blockNum", blockNum, "proposer", proposer)
	timeout := time.After(1000 * time.Millisecond)
	for {
		select {
		case <-timeout:
			return blockProposal, errors.New("block proposal cache timout")
		default:
			// keep querying cache until timeout reached
			blockProposal, err := d.GetBlockProposal(blockNum, proposer)
			if err == nil {
				return blockProposal, nil
			}
		}
	}
}
