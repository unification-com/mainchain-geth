package dsg

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/unification-com/mainchain/common"
	"math/big"
)

const (
	inmMemoryProposals  = 64 // Number of recent block proposals to keep in memory
	inMemoryValidations = 4096 // Number of recent validation messages to keep in memory
)

type Cache struct {
	validations *lru.ARCCache
	proposals *lru.ARCCache
}

type Validation struct {
	BlockNum uint64      `json:"blocknum"`
	Validator uint64     `json:"validator"`
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
	d.proposals.Add(bp.BlockHash, bp)
}


func (d *Cache) InsertValidationMessage(msg ValidationMessage) bool {
	return d.insertValidationMessage(common.NumSignersInRound, *msg.Number, msg.BlockHash, *msg.VerifierId, msg.Authorize)
}

func (d *Cache) insertValidationMessage(totalSigners uint64, blockNumber big.Int, blockHash common.Hash, verifierID big.Int, authorize bool) bool {
	n := blockNumber.Uint64()
	v := verifierID.Uint64()

	key := Validation{
		BlockNum: n,
		Validator: v,
	}

    d.validations.Add(key, authorize)

	return d.acceptBlock(totalSigners, n)
}

func (d *Cache) acceptBlock(totalSigners uint64, blockNum uint64) bool {
	acks := float64(0)

	for i := uint64(0); i < totalSigners; i++ {
		lookupKey := Validation{
			BlockNum: blockNum,
			Validator: i,
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
	accept := acknowledges >= requirement

	return accept
}
