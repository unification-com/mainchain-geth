package dsg

import (
	"github.com/unification-com/mainchain/common"
	"math/big"
)

type Cache struct {
	Validations map[uint64]map[uint64]bool `json:"validations"`
	Proposals []BlockProposal `json:"proposals"`
}

func NewCache() *Cache {
	cache := &Cache{
		Validations: map[uint64]map[uint64]bool{},
		Proposals: make([]BlockProposal, 100),
	}
	return cache
}

func (d *Cache) InsertBlockProposal(bp BlockProposal) {
	d.Proposals = append(d.Proposals, bp)
}


func (d *Cache) InsertValidationMessage(msg ValidationMessage) bool {
	return d.insertValidationMessage(common.ActiveSigners, *msg.Number, msg.BlockHash, *msg.VerifierId, msg.Authorize)
}

func (d *Cache) insertValidationMessage(totalSigners uint64, blockNumber big.Int, blockHash common.Hash, verifierID big.Int, authorize bool) bool {
	n := blockNumber.Uint64()
	v := verifierID.Uint64()

	_, ok := d.Validations[n]
	if ! ok {
		d.Validations[n] = map[uint64]bool{}
	}
	d.Validations[n][v] = authorize

	acks := float64(0)
	for _, value := range d.Validations[n] {
		if value == true {
			acks = acks + 1
		}
	}

	totalSignersFloat := float64(totalSigners)
	requirement := float64(2) / totalSignersFloat
	acknowledges := acks / totalSignersFloat
	accept := acknowledges >= requirement

	return accept
}
