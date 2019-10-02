package dsg

import (
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/log"
)

func (bp *BlockProposal) ValidateBlockProposal(parentHeader *types.Header, cache *Cache) bool {
	numTimeouts := cache.GetInvalidCounter()
	currentSlot := parentHeader.SlotCount + 1 + numTimeouts

	log.Info("ValidateBlockProposal", "bp.Proposer", bp.Proposer)
	expectedSignerIndex, _ := EVSlot(currentSlot)
	actualSignerIndex := EVIdFromEtherbase(bp.Proposer)

	if expectedSignerIndex != actualSignerIndex {
		log.Info("Invalid block proposal - expectedSignerIndex mismatch", "block", bp.Number.Uint64(), "exp", expectedSignerIndex, "got", actualSignerIndex)
		return false
	} else {
		log.Trace("Proposer OK", "id", expectedSignerIndex)
	}

	txHash := bp.ProposedBlock.TxHash()
	txDerivedHash := types.DeriveSha(bp.ProposedBlock.Transactions())

	if txHash != txDerivedHash {
		log.Info("Invalid block proposal - txHash mismatch", "block", bp.Number.Uint64(), "exp", txDerivedHash, "got", txHash)
		return false
	} else {
		log.Trace("txHash OK", "txHash", txHash)
	}

	// TODO: check against this EV's own block TxHash
	// TODO: check bp.Number = parent_height + 1
	return true
}


func ValidateNewBlock(block *types.Block) bool {
	// TODO: Check signatures once they have been added
	return true
}