package dsg

import (
	"bytes"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/consensus"
	"github.com/unification-com/mainchain/core/state"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/log"
	"github.com/unification-com/mainchain/params"
	"github.com/unification-com/mainchain/rpc"
	"math/big"
	"time"
)

var (
	dsgExtraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	dsgExtraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for Sealing EV's signature
)

type Dsg struct {
	config   *params.DsgConfig // Consensus engine configuration parameters
	dsgCache *Cache            // DSG Cache
}

// NewDsg returns a new DSG consensus.Engine
func NewDsg(config *params.DsgConfig) *Dsg {
	return &Dsg{
		config:   config,
		dsgCache: NewCache(),
	}
}

// Author implements engine.Author
func (dsg *Dsg) Author(header *types.Header) (common.Address, error) {
	log.Info("dsg engine.Author", "header.Number", header.Number)
	return common.Address{}, nil
}

// VerifyHeader implements engine.VerifyHeader
func (dsg *Dsg) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	log.Info("dsg engine.VerifyHeader", "header.Number", header.Number)
	return nil
}

// VerifyHeaders implements engine.VerifyHeaders
func (dsg *Dsg) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	log.Info("dsg engine.VerifyHeaders", "CurrentHeader", chain.CurrentHeader().Number)
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	return abort, results
}

// VerifyUncles implements engine.VerifyUncles
func (dsg *Dsg) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	log.Info("dsg engine.VerifyUncles", "CurrentHeader", chain.CurrentHeader().Number)
	return nil
}

// VerifySeal implements engine.VerifySeal
func (dsg *Dsg) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	log.Info("dsg engine.VerifySeal", "header.Number", header.Number)
	return nil
}

// Prepare implements engine.Prepare
func (dsg *Dsg) Prepare(chain consensus.ChainReader, header *types.Header) error {
	log.Info("dsg engine.Prepare", "header.Number", header.Number)

	number := header.Number.Uint64()

	header.Nonce = types.BlockNonce{} // Nonce not used in DSG. Set to empty
	header.Difficulty = CalcDifficulty() // Difficulty not used in DSG. Set to 0x1
	header.Coinbase = common.Address{} // Coinbase not currently used in DSG

	header.MixDigest = common.Hash{} // Todo: set MixDigest to constant hash to identify DSG blocks

	// check this is a valid child block
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// pre-fill initial 32 bytes with empty vanity data if required
	if len(header.Extra) < dsgExtraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, dsgExtraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:dsgExtraVanity]
	// Todo - append RLP encoded empty DsgExtra struct{}

	// Set block's timestamp
	header.Time = parent.Time + dsg.config.BlockTime
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	return nil
}

// Finalize implements engine.Finalize
func (dsg *Dsg) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	log.Info("dsg engine.Finalize", "header.Number", header.Number)

	// Todo - add block rewards to EV before setting header.Root
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Uncles are not used in DSG
	header.UncleHash = types.CalcUncleHash(nil)
}

// FinalizeAndAssemble implements engine.FinalizeAndAssemble
func (dsg *Dsg) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	log.Info("dsg engine.FinalizeAndAssemble", "header.Number", header.Number)

	// Todo - add block rewards to EV before setting header.Root
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Uncles are not used in DSG
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Seal implements engine.Seal
func (dsg *Dsg) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	log.Info("dsg engine.Seal", "CurrentHeader", chain.CurrentHeader().Number)
	return nil
}

// SealHash implements engine.SealHash
func (dsg *Dsg) SealHash(header *types.Header) common.Hash {
	log.Info("dsg engine.SealHash", "header.Number", header.Number)
	return common.Hash{}
}

// CalcDifficulty implements engine.CalcDifficulty
func (dsg *Dsg) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	log.Info("dsg engine.CalcDifficulty", "CurrentHeader", chain.CurrentHeader().Number)
	return CalcDifficulty()
}

// APIs implements engine.APIs
func (dsg *Dsg) APIs(chain consensus.ChainReader) []rpc.API {
	log.Info("dsg engine.APIs")
	return []rpc.API{}
}

// Close implements engine.Close
func (dsg *Dsg) Close() error {
	log.Info("dsg engine.Close")
	return nil
}

func CalcDifficulty() *big.Int {
	return common.Big1
}
