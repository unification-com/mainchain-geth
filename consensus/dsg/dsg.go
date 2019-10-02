package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/log"
	"math/big"
)

// BlockProposal represents a block proposal in DSG.
type BlockProposal struct {
	Number        *big.Int       `json:"number"     gencodec:"required"`
	BlockHash     common.Hash    `json:"blockHash"  gencodec:"required"`
	ProposedBlock *types.Block   `json:"block"      gencodec:"required"`
	Signature     common.Hash    `json:"sig"        gencodec:"required"`
	Proposer      common.Address `json:"proposer"   gencodec:"required"`
}

// ValidationMessage represents a validation message in DSG.
type ValidationMessage struct {
	Number    *big.Int       `json:"number"     gencodec:"required"`
	BlockHash common.Hash    `json:"blockHash"  gencodec:"required"`
	Verifier  common.Address `json:"verifierId" gencodec:"required"`
	Proposer  common.Address `json:"proposerId" gencodec:"required"`
	Signature common.Hash    `json:"signature"  gencodec:"required"`
	Authorize bool           `json:"authorize"  gencodec:"required"`
}

// RequestNewBlockProposalMessage is used when a new BP is required
type RequestNewBlockProposalMessage struct {
	Number    *big.Int       `json:"number"     gencodec:"required"`
	Verifier  common.Address `json:"verifierId" gencodec:"required"`
	Signature common.Hash    `json:"signature"  gencodec:"required"`
}

func ProposeBlock(block *types.Block, proposer common.Address) BlockProposal {

	log.Info("Propose block #", "num", block.Number().String(), "proposer", proposer)

	proposedBlock := BlockProposal{
		Number:        block.Number(),
		BlockHash:     block.Hash(),
		Proposer:      proposer,
		ProposedBlock: block,
		Signature:     common.Hash{}, // TODO - sign
	}

	return proposedBlock
}

func Authorized(parentHeader types.Header, numInvalids uint64, etherbase common.Address) bool {
	slot := parentHeader.SlotCount + 1 + numInvalids

	expectedSignerIndex, _ := EVSlot(slot)
	actualSignerIndex := EVIdFromEtherbase(etherbase)

	return expectedSignerIndex == actualSignerIndex
}

func SetSlotNumber(parentHeader types.Header, block *types.Block, numInvalids uint64) *types.Block {
	// if parent was genesis, use 0 as parentSlotCount

	parentSlotCount := uint64(0)
	if block.Number().Cmp(big.NewInt(1)) == 1 {
		parentSlotCount = parentHeader.SlotCount
	}

	header := block.Header()
	header.SlotCount = parentSlotCount + 1 + numInvalids
	newBlock := block.WithSeal(header)

	return newBlock
}

func EVSlotInternal(slotNumber uint64, blocksInEpoch uint64, numRounds uint64, numSigners uint64) (uint64, uint64) {
	var slotNumberBase0 = slotNumber - 1
	var slotIndex = slotNumberBase0 % blocksInEpoch
	var epochNumber = slotNumberBase0 / blocksInEpoch

	var quadrant = slotIndex / (blocksInEpoch / numRounds)
	var numSignersInQuadrant = numSigners / numRounds
	var factor = numSigners / numRounds
	var position = slotIndex % numSignersInQuadrant
	var signerIndex = quadrant*factor + (position + 1)

	return signerIndex, epochNumber
}

// The base 0 signer index for a given slot number
// where the genesis block is slot 0, and the current Epoch
func EVSlot(slotNumber uint64) (uint64, uint64) {
	return EVSlotInternal(slotNumber, common.BlocksInEpoch, common.NumberOfRounds, common.NumSignersInRound)
}

// F is the fault function to calculate the number of required ACKs/NACKs
func F() (float64, float64) {
	ackRequirement := float64(2) / float64(3)
	nackRequirement := float64(1) / float64(3)
	return ackRequirement, nackRequirement
}
