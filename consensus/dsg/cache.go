package dsg

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/log"
)

const (
	inmMemoryProposals  = 96 // Number of recent block proposals to keep in memory
	inMemoryValidations = 4096 // Number of recent validation messages to keep in memory
)

type Cache struct {
	validations *lru.ARCCache
	proposals *lru.ARCCache
}

type ValidationKey struct {
	BlockNum uint64      `json:"blocknum"`
	Proposer uint64      `json:"proposer"`
	Validator uint64     `json:"validator"`
}

type ProposalKey struct {
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

	key := ProposalKey{
		BlockNum: n,
		Proposer: p,
	}

	d.proposals.Add(key, bp)
}


func (d *Cache) InsertValidationMessage(msg ValidationMessage) bool {
	return d.insertValidationMessage(msg)
}

func (d *Cache) insertValidationMessage(msg ValidationMessage) bool {
	n := msg.Number.Uint64()
	v := msg.VerifierId.Uint64()
	p := msg.ProposerId.Uint64()

	key := ValidationKey{
		BlockNum: n,
		Validator: v,
		Proposer: p,
	}

	log.Info("insertValidationMessage", "block", key.BlockNum, "proposer", key.Proposer, "validator", key.Validator, "authorise", msg.Authorize)
    d.validations.Add(key, msg)

	return d.acceptBlock(n, p)
}

func (d *Cache) acceptBlock(blockNum uint64, proposerId uint64) bool {
	acks := float64(0)

	for i := uint64(0); i < common.NumSignersInRound; i++ {
		lookupKey := ValidationKey{
			BlockNum: blockNum,
			Validator: i,
			Proposer: proposerId,
		}

		var validationMessage ValidationMessage
		if vm, ok := d.validations.Get(lookupKey); ok {
			if validationMessage, ok = vm.(ValidationMessage); ok {
				if validationMessage.Authorize == true {
					acks = acks + 1
				}
			}
		}
	}

	totalSignersFloat := float64(common.NumSignersInRound)
	requirement := float64(2) / totalSignersFloat
	acknowledges := acks / totalSignersFloat
	log.Info("acceptBlock", "num", blockNum, "proposer", proposerId, "acks", acks)
	return acknowledges >= requirement
}
