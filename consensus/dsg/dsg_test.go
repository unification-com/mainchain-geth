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
	type EVSlotTestCase struct {
		SlotNumber uint64
		SignerIndex uint64
		EpochNumber uint64
	}

	testCases := []EVSlotTestCase{
		{1, 1, 0},
		{2, 2, 0},
		{3, 3, 0},
		{4, 1, 0},
		{5, 2, 0},
		{6, 3, 0},
		{7, 4, 0},
		{8, 5, 0},
		{9, 6, 0},
		{10, 4, 0},
		{11, 5, 0},
		{12, 6, 0},
		{13, 7, 0},
		{14, 8, 0},
		{15, 9, 0},
		{16, 7, 0},
		{17, 8, 0},
		{18, 9, 0},
		{19, 10, 0},
		{20, 11, 0},
		{21, 12, 0},
		{22, 10, 0},
		{23, 11, 0},
		{24, 12, 0},
		{25, 1, 1},
		{26, 2, 1},
		{27, 3, 1},
		{28, 1, 1},
		{29, 2, 1},
		{30, 3, 1},
	}

	for _, testCase := range testCases {
		baseZeroSignerIndex, epochNumber := EVSlotInternal(testCase.SlotNumber, 24, 4, 12)
		baseOneSignerIndex := baseZeroSignerIndex + 1
		assertEqual(t, baseOneSignerIndex, testCase.SignerIndex)
		assertEqual(t, epochNumber, testCase.EpochNumber)
	}
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
