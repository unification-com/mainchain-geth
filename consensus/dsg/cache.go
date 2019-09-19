package dsg

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/log"
	"math/big"
)

const (
	inmMemoryProposals  = 96 // Number of recent block proposals to keep in memory
	inMemoryValidations = 4096 // Number of recent validation messages to keep in memory
)

type Cache struct {
	validations *lru.ARCCache
	proposals *lru.ARCCache
}

type Validation struct {
	BlockNum uint64      `json:"blocknum"`
	Proposer uint64      `json:"proposer"`
	Validator uint64     `json:"validator"`
}

type Proposal struct {
	BlockNum uint64      `json:"blocknum"`
	Proposer uint64      `json:"proposer"`
}

func NewCache() *Cache {

	validations, _ := lru.NewARC(inMemoryValidations)
	proposals, _ := lru.NewARC(inmMemoryProposals)

	cache := &Cache{
		validations: validations,
		proposals: proposals,
	}
	return cache
}

func (d *Cache) InsertBlockProposal(bp BlockProposal) {
	n := bp.Number.Uint64()
	p := bp.ProposerId.Uint64()

	key := Proposal{
		BlockNum: n,
		Proposer: p,
	}

	d.proposals.Add(key, bp)
}


func (d *Cache) InsertValidationMessage(msg ValidationMessage) bool {
	return d.insertValidationMessage(common.NumSignersInRound, *msg.Number, msg.BlockHash, *msg.VerifierId, *msg.ProposerId, msg.Authorize)
}

func (d *Cache) insertValidationMessage(totalSigners uint64, blockNumber big.Int, blockHash common.Hash, verifierID big.Int, proposerId big.Int, authorize bool) bool {
	n := blockNumber.Uint64()
	v := verifierID.Uint64()
	p := proposerId.Uint64()

	log.Info("insertValidationMessage", "block", n, "proposer", p, "validator", v, "authorise", authorize)

	key := Validation{
		BlockNum: n,
		Validator: v,
		Proposer: p,
	}

    d.validations.Add(key, authorize)

	return d.acceptBlock(totalSigners, n, p)
}

func (d *Cache) acceptBlock(totalSigners uint64, blockNum uint64, proposerId uint64) bool {
	acks := float64(0)

	for i := uint64(0); i < totalSigners; i++ {
		lookupKey := Validation{
			BlockNum: blockNum,
			Validator: i,
			Proposer: proposerId,
		}
		if a, ok := d.validations.Get(lookupKey); ok {
			if a == true {
				acks = acks + 1
			}
		}
	}

	totalSignersFloat := float64(totalSigners)
	requirement := float64(2) / totalSignersFloat
	acknowledges := acks / totalSignersFloat
	log.Info("acceptBlock", "num", blockNum, "proposer", proposerId, "acks", acks)
	return acknowledges >= requirement

}
