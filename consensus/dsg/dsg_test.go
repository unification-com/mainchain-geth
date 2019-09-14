package dsg

import (
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
