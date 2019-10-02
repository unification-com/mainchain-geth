package dsg

import (
	"github.com/unification-com/mainchain/common"
	"math/big"
	"testing"
)

type tuple struct {
	Verifier string
	Authorize bool
}

func TestCacheAccept(t *testing.T) {
	hash := common.BigToHash(big.NewInt(1))
	proposer := common.HexToAddress("1")

	cache := NewCache()

	cases := []tuple{
		{"1", true},
		{"2", true},
		{"3", false},
	}

	result := Unknown
	for _, testCase := range cases {
		msg := ValidationMessage{Number: big.NewInt(1), BlockHash: hash, Verifier: common.HexToAddress(testCase.Verifier), Proposer: proposer, Authorize:testCase.Authorize}
		result = cache.insertValidationMessage(msg)
	}
	assertEqual(t, result, Accept)
}

func TestCacheReject(t *testing.T) {
	hash := common.BigToHash(big.NewInt(1))
	proposer := common.HexToAddress("1")

	cache := NewCache()

	cases := []tuple{
		{"1", false},
		{"2", false},
	}

	result := Unknown
	for _, testCase := range cases {
		msg := ValidationMessage{Number: big.NewInt(1), BlockHash: hash, Verifier: common.HexToAddress(testCase.Verifier), Proposer: proposer, Authorize:testCase.Authorize}
		result = cache.insertValidationMessage(msg)
	}
	assertEqual(t, result, Reject)
}

func TestCacheUnknown(t *testing.T) {
	hash := common.BigToHash(big.NewInt(1))
	proposer := common.HexToAddress("1")

	cache := NewCache()

	cases := []tuple{
		{"1", true},
	}

	result := Unknown
	for _, testCase := range cases {
		msg := ValidationMessage{Number: big.NewInt(1), BlockHash: hash, Verifier: common.HexToAddress(testCase.Verifier), Proposer: proposer, Authorize:testCase.Authorize}
		result = cache.insertValidationMessage(msg)
	}
	assertEqual(t, result, Unknown)
}
