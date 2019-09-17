package dsg

func (bp *BlockProposal) ValidateBlockProposal() bool {
	signerIndex, _ := EVSlot(bp.Number.Uint64())
	return signerIndex == bp.ProposerId.Uint64()
}
