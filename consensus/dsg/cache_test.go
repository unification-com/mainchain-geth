package dsg

import (
	"github.com/unification-com/mainchain/common"
	"math/big"
	"testing"
)

func TestCache(t *testing.T) {
	hash := common.BigToHash(big.NewInt(1))
	proposer := common.HexToAddress("1")

	cache := NewCache()

	accept1 := cache.insertValidationMessage(ValidationMessage{Number: big.NewInt(1), BlockHash: hash, Verifier: common.HexToAddress("1"), Proposer: proposer, Authorize: false})
	accept2 := cache.insertValidationMessage(ValidationMessage{Number: big.NewInt(1), BlockHash: hash, Verifier: common.HexToAddress("2"), Proposer: proposer, Authorize: true})
	accept3 := cache.insertValidationMessage(ValidationMessage{Number: big.NewInt(1), BlockHash: hash, Verifier: common.HexToAddress("3"), Proposer: proposer, Authorize: true})

	assertEqual(t, accept1, false)
	assertEqual(t, accept2, false)
	assertEqual(t, accept3, true)
}
