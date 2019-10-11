package dsg

import (
	"fmt"
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
		baseOneSignerIndex, epochNumber := EVSlotInternal(testCase.SlotNumber)
		assertEqual(t, baseOneSignerIndex, testCase.SignerIndex)
		assertEqual(t, epochNumber, testCase.EpochNumber)
	}
}


func TestEVSetForRound(t *testing.T) {
	type EVSetTestCase struct {
		SlotNumber uint64
		Signers []uint64
	}

	testCases := []EVSetTestCase{
		{1, []uint64{1, 2, 3}},
		{2, []uint64{1, 2, 3}},
		{3, []uint64{1, 2, 3}},
		{4, []uint64{1, 2, 3}},
		{5, []uint64{1, 2, 3}},
		{6, []uint64{1, 2, 3}},
		{7, []uint64{4, 5, 6}},
		{8, []uint64{4, 5, 6}},
		{9, []uint64{4, 5, 6}},
		{10, []uint64{4, 5, 6}},
		{11, []uint64{4, 5, 6}},
		{12, []uint64{4, 5, 6}},
		{13, []uint64{7, 8, 9}},
	}
	for _, testCase := range testCases {
		signers := EVSetInternal(testCase.SlotNumber, 24, 4, 12)
		fmt.Printf("signers: %v\n", signers)
	}
}

func TestGetValidationPool(t *testing.T) {
	validationPool := GetValidatorPool()

	for i, vp := range validationPool {
		t.Log("validationPool", i, vp.Hex())
	}
}

func TestEVIdFromEtherbase(t *testing.T) {
	assertEqual(t, EVIdFromEtherbase(common.HexToAddress("0x001A320943d4535e93d31E4A65a6e21C5dF375D7")), uint64(1))
	assertEqual(t, EtherbaseFromEVId(1), common.HexToAddress("0x001A320943d4535e93d31E4A65a6e21C5dF375D7"))

	assertEqual(t, EVIdFromEtherbase(common.HexToAddress("0x004A435F1D54aA5cc9FCfA0fEB6B8c4a428bbB93")), uint64(4))
	assertEqual(t, EtherbaseFromEVId(4), common.HexToAddress("0x004A435F1D54aA5cc9FCfA0fEB6B8c4a428bbB93"))
}
