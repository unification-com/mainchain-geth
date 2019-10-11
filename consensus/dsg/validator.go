package dsg

import (
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/log"
)

func (d *Dsg) ValidateBlockProposal(bp BlockProposal) bool {

	currentSlot := d.calculateSlot(bp.ProposedBlock.Header())

	log.Info("ValidateBlockProposal", "bp.Proposer", bp.Proposer)
	expectedSignerIndex, _ := d.EVSlot(currentSlot)
	actualSignerIndex := d.EVIdFromEtherbase(bp.Proposer)

	if expectedSignerIndex != actualSignerIndex {
		log.Info("Invalid block proposal - expectedSignerIndex mismatch", "bpnum", bp.Number, "numInv", d.dsgCache.GetInvalidCounter(), "slot", currentSlot, "exp", expectedSignerIndex, "got", actualSignerIndex)
		return false
	}

	txHash := bp.ProposedBlock.TxHash()
	txDerivedHash := types.DeriveSha(bp.ProposedBlock.Transactions())

	if txHash != txDerivedHash {
		log.Info("Invalid block proposal - txHash mismatch", "block", bp.Number, "txs", len(bp.ProposedBlock.Transactions()), "exp", txDerivedHash, "got", txHash)
		return false
	}

	// TODO: check against this EV's own block TxHash
	// TODO: check bp.Number = parent_height + 1
	return true
}


func (d *Dsg) ValidateNewBlock(block *types.Block) bool {
	// TODO: Check signatures once they have been added
	return true
}