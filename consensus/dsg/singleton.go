package dsg

import (
	"bytes"
	"github.com/unification-com/mainchain/accounts"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/rlp"
	"io"
	"sync"
)

// backing account.
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

// SignerFn is a signer callback function to request a header to be signed by a

type Singleton struct {
	signer common.Address // Ethereum address of the signing key
	signFn SignerFn       // Signer function to authorize hashes with
	lock   sync.RWMutex   // Protects the signer fields
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (c *Singleton) Authorize(signer common.Address, signFn SignerFn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.signer = signer
	c.signFn = signFn
}

func (c *Singleton) Sign(block *types.Block) *types.Block {
	header := block.Header()
	sighash, _ := c.signFn(accounts.Account{Address: c.signer}, accounts.MimetypeClique, DSGRLP(header))
	extraSeal := 65
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)

	newblock := block.WithSeal(header)
	return newblock
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

func DSGRLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)
	return b.Bytes()
}
