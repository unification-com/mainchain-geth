package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/consensus"
	"github.com/unification-com/mainchain/core/state"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/params"
	"github.com/unification-com/mainchain/rpc"
	"math/big"
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
	return common.Address{}, nil
}

// VerifyHeader implements engine.VerifyHeader
func (dsg *Dsg) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

// VerifyHeaders implements engine.VerifyHeaders
func (dsg *Dsg) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	return abort, results
}

// VerifyUncles implements engine.VerifyUncles
func (dsg *Dsg) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	return nil
}

// VerifySeal implements engine.VerifySeal
func (dsg *Dsg) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

// Prepare implements engine.Prepare
func (dsg *Dsg) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

// Finalize implements engine.Finalize
func (dsg *Dsg) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {

}

// FinalizeAndAssemble implements engine.FinalizeAndAssemble
func (dsg *Dsg) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Seal implements engine.Seal
func (dsg *Dsg) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	return nil
}

// SealHash implements engine.SealHash
func (dsg *Dsg) SealHash(header *types.Header) common.Hash {
	return common.Hash{}
}

// CalcDifficulty implements engine.CalcDifficulty
func (dsg *Dsg) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	return big.NewInt(0)
}

// APIs implements engine.APIs
func (dsg *Dsg) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{}
}

// Close implements engine.Close
func (dsg *Dsg) Close() error {
	return nil
}
