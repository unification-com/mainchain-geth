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

type VerifiedBlock struct {
	FinalBlock *types.Block
}

func (d *Dsg) ProposeBlock(block *types.Block, proposer common.Address) BlockProposal {

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

// Authorized checks if theEV is authorised to propose this block
// Takes in header of the block to be proposed
func (d *Dsg) Authorized(header *types.Header) bool {
	number := header.Number.Uint64()
	parentHeader := d.chain.GetHeader(header.ParentHash, number-1)

	numInvalids := d.dsgCache.GetInvalidCounter()
	slot := parentHeader.SlotCount + 1 + numInvalids

	expectedSignerIndex, _ := d.EVSlot(slot)
	actualSignerIndex := d.EVIdFromEtherbase(d.signer)

	log.Info("dsg authorise", "block", number, "parent", parentHeader.Number, "numInvalids", numInvalids, "p_slot", parentHeader.SlotCount, "slot", slot, "etherbase", d.signer, "exp", expectedSignerIndex, "act", actualSignerIndex)

	return expectedSignerIndex == actualSignerIndex
}

func (d *Dsg) SetSlotNumberInBlock(block *types.Block) *types.Block {
	slot := d.calculateSlot(block.Header())

	header := block.Header()
	header.SlotCount = slot
	// Todo - need to re-sign extraData here?
	newBlock := block.WithSeal(header)

	return newBlock
}

func (d *Dsg) EVSlotInternal(slotNumber uint64) (uint64, uint64) {
	blocksInEpoch := d.config.BlocksInEpoch
	numRounds := d.config.NumberOfRounds
	numSigners := d.config.NumSignersinRound

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

func (d *Dsg) contains(a uint64, list []uint64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (d *Dsg) EVSetInternal(slotNumber uint64) []uint64 {

	blocksInEpoch := d.config.BlocksInEpoch
	numRounds := d.config.NumberOfRounds
	numSigners := d.config.NumSignersinRound

	var slotNumberBase0 = slotNumber - 1
	var slotIndex = slotNumberBase0 % blocksInEpoch

	var quadrant = slotIndex / (blocksInEpoch / numRounds)
	var numSignersInQuadrant = numSigners / numRounds
	var factor = numSigners / numRounds

	var signerIndex = quadrant*factor
	var ret []uint64

	for i := 0; i < int(numSignersInQuadrant); i++ {
		inc := signerIndex + uint64(1) + uint64(i)
		ret = append(ret, inc)
	}
	return ret
}

func (d *Dsg) EVSet(slotNumber uint64) []uint64 {
	return d.EVSetInternal(slotNumber)
}

func (d *Dsg) Member(slotNumber uint64, address common.Address) bool {
	set := d.EVSet(slotNumber)
	evID := d.EVIdFromEtherbase(address)
	return d.contains(evID, set)
}

// The base 0 signer index for a given slot number
// where the genesis block is slot 0, and the current Epoch
func (d *Dsg) EVSlot(slotNumber uint64) (uint64, uint64) {
	return d.EVSlotInternal(slotNumber)
}

// calculateSlot returns the slot for the current block being processed
// by the network. Slot is based on parent.SlotCount + 1 + numInvalids
func (d *Dsg) calculateSlot(header *types.Header) uint64 {
	slot := uint64(1)
	number := header.Number.Uint64()
	parent := d.chain.GetHeader(header.ParentHash, number-1)
	if parent != nil {
		slot = parent.SlotCount + d.dsgCache.GetInvalidCounter() + 1
	} else {
		log.Info("calculateSlot - parent not found", "num", number, "parentHash", header.ParentHash)
		slot = d.dsgCache.GetInvalidCounter() + 1
	}

	return slot
}

// F is the fault function to calculate the number of required ACKs/NACKs
func F() (float64, float64) {
	ackRequirement := float64(2) / float64(3)
	nackRequirement := float64(1) / float64(3)
	return ackRequirement, nackRequirement
}
