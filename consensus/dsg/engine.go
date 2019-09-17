package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/consensus"
	"github.com/unification-com/mainchain/core/state"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/rpc"
	"math/big"
)

type Dsg struct {
	config *Config // Consensus engine configuration parameters
}

// Author implements engine.Author
func (c *Dsg) Author(header *types.Header) (common.Address, error) {
	return common.Address{}, nil
}

// VerifyHeader implements engine.VerifyHeader
func (c *Dsg) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

// VerifyHeaders implements engine.VerifyHeaders
func (c *Dsg) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	return abort, results
}

// VerifyUncles implements engine.VerifyUncles
func (c *Dsg) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	return nil
}

// VerifySeal implements engine.VerifySeal
func (c *Dsg) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

// Prepare implements engine.Prepare
func (c *Dsg) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

// Finalize implements engine.Finalize
func (c *Dsg) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {

}

// FinalizeAndAssemble implements engine.FinalizeAndAssemble
func (c *Dsg) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Seal implements engine.Seal
func (c *Dsg) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	return nil
}

// SealHash implements engine.SealHash
func (c *Dsg) SealHash(header *types.Header) common.Hash {
	return common.Hash{}
}

// CalcDifficulty implements engine.CalcDifficulty
func (c *Dsg) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	return big.NewInt(0)
}

// APIs implements engine.APIs
func (c *Dsg) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{}
}

// Close implements engine.Close
func (c *Dsg) Close() error {
	return nil
}
