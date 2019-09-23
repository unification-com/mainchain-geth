package dsg

import (
	"github.com/unification-com/mainchain/common"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestEVSlots(t *testing.T) {
	signerIndex, epochNumber := EVSlotInternal(10, 24, 4, 12)
	assertEqual(t, signerIndex, uint64(3))
	assertEqual(t, epochNumber, uint64(0))
}

func TestGetValidationPool(t *testing.T) {
	validationPool := GetValidatorPool()

	for i, vp := range validationPool {
		t.Log("validationPool", i, vp.Hex())
	}
}

func TestEVIdFromEtherbase(t *testing.T) {
	etherbase := common.HexToAddress("0x004A435F1D54aA5cc9FCfA0fEB6B8c4a428bbB93")
	evID := EVIdFromEtherbase(etherbase)
	assertEqual(t, evID, uint64(3))
}
