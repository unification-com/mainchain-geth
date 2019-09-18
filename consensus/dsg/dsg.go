package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/state"
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
	ProposerId    *big.Int       `json:"proposerid" gencodec:"required"`
	ProposedBlock *types.Block   `json:"block"      gencodec:"required"`
	Signature     common.Hash    `json:"sig"        gencodec:"required"`
	Address       common.Address `json:"address"    gencodec:"required"`
}

// ValidationMessage represents a validation message in DSG.
type ValidationMessage struct {
	Number     *big.Int    `json:"number"     gencodec:"required"`
	BlockHash  common.Hash `json:"blockHash"  gencodec:"required"`
	VerifierId *big.Int    `json:"verifierId" gencodec:"required"`
	Authorize  bool        `json:"authorize"  gencodec:"required"`
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

func ProposeBlock(block *types.Block, proposer common.Address) error {

	log.Info("Propose block #", "num", block.Number().String())

	proposedBlock := BlockProposal{
		Number:        block.Number(),
		BlockHash:     block.Hash(),
		ProposerId:    getTurn(block.Number()),
		ProposedBlock: block,
		Signature:     common.Hash{}, // TODO - sign
		Address:       proposer,
	}

	encodedBlockProposal, err := rlp.EncodeToBytes(proposedBlock)

	if err != nil {
		log.Error("ProposeBlock encode error:", err)
	}

	log.Info("encodedBlockProposal", "enc", encodedBlockProposal)

	return nil
}

func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

func getTurn(blockNumber *big.Int) *big.Int {
	d := blockNumber.Sub(blockNumber, big.NewInt(1))
	turn := blockNumber.Mod(d, big.NewInt(common.NumSignersInRound))
	return turn
}

func valid(statedb *state.StateDB, blockNumber *big.Int, signer common.Address) bool {
	turn := getTurn(blockNumber)

	var whitelist []common.Address

	for i := 0; i < common.NumSignersInRound; i++ {
		keyhash := statedb.GetState(common.HexToAddress(common.DSG), common.BigToHash(big.NewInt(int64(i))))
		whitelist = append(whitelist, common.BytesToAddress(keyhash[:]))
	}

	allowed := whitelist[turn.Int64()]

	return allowed == signer
}

func DSGSeal(statedb *state.StateDB, block *types.Block, signer common.Address) (bool, error) {
	v := valid(statedb, block.Number(), signer)

	return v, nil
}

func DSGFilter(statedb *state.StateDB, block types.Block) (bool, error) {
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
	v := valid(statedb, block.Number(), signer)
	return v, nil
}

func EVSlotInternal(blockNumber uint64, blocksInEpoch uint64, numQuarters uint64, activeSigners uint64) (uint64, uint64) {
	var blockNumberBase0 = blockNumber - 1
	var blockIndex = blockNumberBase0 % blocksInEpoch
	var epochNumber = blockNumberBase0 / blocksInEpoch

	var quadrant = blockIndex / (blocksInEpoch / numQuarters)
	var numSignersInQuadrant = activeSigners / numQuarters
	var factor = activeSigners / numQuarters
	var position = blockIndex % numSignersInQuadrant
	var signerIndex = quadrant*factor + (position + 1)

	return signerIndex - 1, epochNumber
}

// The base 0 signer index for a given block number
// where the genesis block is block 0, and the current Epoch
func EVSlot(blockNumber uint64) (uint64, uint64) {
	return EVSlotInternal(blockNumber, common.BlocksInEpoch, common.NumberOfRounds, common.NumSignersInRound)
}

func VerifierIdFromEtherbase(etherbase common.Address) *big.Int {
	// TODO: Implement me
	return big.NewInt(1)
}
