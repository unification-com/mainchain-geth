package dsg

import (
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/log"
)

func (bp *BlockProposal) ValidateBlockProposal() bool {

	expectedSignerIndex, _ := EVSlot(bp.Number.Uint64())
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
