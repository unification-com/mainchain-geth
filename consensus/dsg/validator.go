package dsg

import (
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/log"
)

func (bp *BlockProposal) ValidateBlockProposal() bool {

	signerIndex, _ := EVSlot(bp.Number.Uint64())

	if signerIndex != bp.ProposerId.Uint64() {
		log.Error("Invalid block proposal - signerIndex mismatch", "exp", signerIndex, "got", bp.ProposerId.Uint64())
		return false
	} else {
		log.Trace("ProposerId OK", "id", signerIndex)
	}

	txHash := bp.ProposedBlock.TxHash()
	txDerivedHash := types.DeriveSha(bp.ProposedBlock.Transactions())

	if txHash != txDerivedHash {
		log.Error("Invalid block proposal - txHash mismatch", "exp", txDerivedHash, "got", txHash)
		return false
	} else {
		log.Trace("txHash OK", "txHash", txHash)
	}

	// TODO: check against this EV's own block TxHash
	// TODO: check bp.Number = parent_height + 1
	return true
}
