package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/state"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/rlp"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
)

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

func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

func valid(statedb *state.StateDB, blockNumber *big.Int, signer common.Address) bool {
	d := blockNumber.Sub(blockNumber, big.NewInt(1))
	turn := blockNumber.Mod(d, big.NewInt(3))

	key1hash := statedb.GetState(common.HexToAddress(common.DSG), common.HexToHash("0x0"))
	key2hash := statedb.GetState(common.HexToAddress(common.DSG), common.HexToHash("0x1"))
	key3hash := statedb.GetState(common.HexToAddress(common.DSG), common.HexToHash("0x2"))

	whitelist := []common.Address{
		common.BytesToAddress(key1hash[:]),
		common.BytesToAddress(key2hash[:]),
		common.BytesToAddress(key3hash[:]),
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