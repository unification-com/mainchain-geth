package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/log"
	"github.com/unification-com/mainchain/rlp"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
)

// BlockProposal represents a block proposal in DSG.
type BlockProposal struct {
	Number        *big.Int       `json:"number"     gencodec:"required"`
	BlockHash     common.Hash    `json:"blockHash"  gencodec:"required"`
	ProposedBlock *types.Block   `json:"block"      gencodec:"required"`
	Signature     common.Hash    `json:"sig"        gencodec:"required"`
	Proposer      common.Address `json:"proposer"    gencodec:"required"`
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

func encodeSigHeader(w io.Writer, header *types.Header) {
	err := rlp.Encode(w, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}

func ProposeBlock(block *types.Block, proposer common.Address) BlockProposal {

	log.Info("Propose block #", "num", block.Number().String())

	proposedBlock := BlockProposal{
		Number:        block.Number(),
		BlockHash:     block.Hash(),
		Proposer:      proposer,
		ProposedBlock: block,
		Signature:     common.Hash{}, // TODO - sign
	}

	return proposedBlock
}

func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

func valid(blockNumber uint64, signer common.Address) bool {
	turn, _ := EVSlot(blockNumber)

	whitelist := getStakedWallets()

	allowed := whitelist[turn].Address

	return allowed == signer
}

func DSGSeal(block *types.Block, signer common.Address) (bool, error) {
	v := valid(block.Number().Uint64(), signer)

	return v, nil
}

func DSGFilter(block types.Block) (bool, error) {
	extraSeal := 65
	header := block.Header()
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return false, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
	v := valid(block.Number().Uint64(), signer)
	return v, nil
}

func EVSlotInternal(blockNumber uint64, blocksInEpoch uint64, numRounds uint64, numSigners uint64) (uint64, uint64) {
	var blockNumberBase0 = blockNumber - 1
	var blockIndex = blockNumberBase0 % blocksInEpoch
	var epochNumber = blockNumberBase0 / blocksInEpoch

	var quadrant = blockIndex / (blocksInEpoch / numRounds)
	var numSignersInQuadrant = numSigners / numRounds
	var factor = numSigners / numRounds
	var position = blockIndex % numSignersInQuadrant
	var signerIndex = quadrant*factor + (position + 1)

	return signerIndex - 1, epochNumber
}

// The base 0 signer index for a given block number
// where the genesis block is block 0, and the current Epoch
func EVSlot(blockNumber uint64) (uint64, uint64) {
	return EVSlotInternal(blockNumber, common.BlocksInEpoch, common.NumberOfRounds, common.NumSignersInRound)
}

func ListenForBlockProposal(cache *Cache, blockNum *big.Int, verifier common.Address) ValidationMessage {

	expectedProposer, _ := EVSlot(blockNum.Uint64())
	stakers := getStakedWallets()
	proposerAddress := stakers[expectedProposer].Address
	valid := false

	bp, err := cache.PollBlockProposalCache(blockNum, proposerAddress)

	if err != nil {
		log.Info("listen for block proposal error","err", err)
	} else {
		log.Info("listen success - found bp in cache and validating", "block", bp.Number, "proposer", bp.Proposer)
		valid = bp.ValidateBlockProposal()
	}
	vm := ValidationMessage{Number: bp.Number, BlockHash: bp.BlockHash, Verifier: verifier, Proposer: bp.Proposer, Signature: common.Hash{}, Authorize:valid}
	return vm
}
