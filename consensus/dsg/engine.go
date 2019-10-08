package dsg

import (
	"bytes"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/consensus"
	"github.com/unification-com/mainchain/core/state"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/log"
	"github.com/unification-com/mainchain/params"
	"github.com/unification-com/mainchain/rlp"
	"github.com/unification-com/mainchain/rpc"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
	"time"
)

var (
	defaultBlockTime = 5 // default block time if none is set in DSG Config
	dsgExtraVanity   = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	dsgExtraSeal     = 65 // Fixed number of extra-data suffix bytes reserved for Sealing EV's signature
)

type Dsg struct {
	config   *params.DsgConfig // Consensus engine configuration parameters
	dsgCache *Cache            // DSG Cache
}

// NewDsg returns a new DSG consensus.Engine
func NewDsg(config *params.DsgConfig) *Dsg {

	conf := *config
	if conf.BlockTime == 0 {
		conf.BlockTime = uint64(defaultBlockTime)
	}

	return &Dsg{
		config:   &conf,
		dsgCache: NewCache(),
	}
}

// Author implements engine.Author
func (d *Dsg) Author(header *types.Header) (common.Address, error) {
	log.Info("engine.Author", "header.Number", header.Number)
	return common.Address{}, nil
}

// VerifyHeader implements engine.VerifyHeader
func (d *Dsg) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	log.Info("engine.VerifyHeader", "header.Number", header.Number)
	return nil
}

// VerifyHeaders implements engine.VerifyHeaders
func (d *Dsg) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	log.Info("engine.VerifyHeaders", "CurrentHeader", chain.CurrentHeader().Number)
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	return abort, results
}

// VerifyUncles implements engine.VerifyUncles
func (d *Dsg) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	log.Info("engine.VerifyUncles", "CurrentHeader", chain.CurrentHeader().Number)
	return nil
}

// VerifySeal implements engine.VerifySeal
func (d *Dsg) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	log.Info("engine.VerifySeal", "header.Number", header.Number)
	return nil
}

// Prepare implements engine.Prepare
func (d *Dsg) Prepare(chain consensus.ChainReader, header *types.Header) error {
	log.Info("engine.Prepare", "header.Number", header.Number)

	number := header.Number.Uint64()

	header.Nonce = types.BlockNonce{} // Nonce not used in DSG. Set to empty
	header.Difficulty = CalcDifficulty() // Difficulty not used in DSG. Set to 0x1
	header.Coinbase = common.Address{} // Coinbase not currently used in DSG. Todo: Will be required for block rewards

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
	header.Time = parent.Time + d.config.BlockTime
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	return nil
}

// Finalize implements engine.Finalize
func (d *Dsg) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	log.Info("engine.Finalize", "header.Number", header.Number)

	// Todo - add block rewards to EV before setting header.Root - will need header.Coinbase for this
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Uncles are not used in DSG
	header.UncleHash = types.CalcUncleHash(nil)
}

// FinalizeAndAssemble implements engine.FinalizeAndAssemble
func (d *Dsg) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	log.Info("engine.FinalizeAndAssemble", "header.Number", header.Number)

	// Todo - add block rewards to EV before setting header.Root - will need header.Coinbase for this
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Uncles are not used in DSG
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Seal implements engine.Seal
func (d *Dsg) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	log.Info("engine.Seal", "block.Header()", block.Header().Number)

	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errCannotSealGenesis
	}

	// return error if we're not authorised to propose & seal
	slot := d.getSlot(chain, header)
	member := Member(slot, common.Address{})
	if !member {
		return errNotAuthorisedToPropose
	}

	// set delay to specified blocktime
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now())
	select {
	case <-stop:
		return nil
	case <-time.After(delay):
	}

	// Todo - kick off process to propose, validate etc.
	// Todo - Need to pull in all message handling, channels, tasks etc. internally
	// Todo - push all signatures into extraData
	// Todo - final verified block.WithSeal(header) needs to be pushed to the results chan<-

	return nil
}

// SealHash implements engine.SealHash
func (d *Dsg) SealHash(header *types.Header) common.Hash {
	log.Info("engine.SealHash", "header.Number", header.Number)
	return SealHash(header)
}

// CalcDifficulty implements engine.CalcDifficulty
func (d *Dsg) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	log.Info("engine.CalcDifficulty", "CurrentHeader", chain.CurrentHeader().Number)
	return CalcDifficulty()
}

// APIs implements engine.APIs
func (d *Dsg) APIs(chain consensus.ChainReader) []rpc.API {
	log.Info("engine.APIs")
	return []rpc.API{}
}

// Close implements engine.Close
func (d *Dsg) Close() error {
	log.Info("engine.Close")
	return nil
}

func CalcDifficulty() *big.Int {
	return common.Big1
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

func (d *Dsg) getSlot(chain consensus.ChainReader, header *types.Header) uint64 {
	number := header.Number.Uint64()
	parent := chain.GetHeader(header.ParentHash, number-1)
	slot := parent.SlotCount + d.dsgCache.GetInvalidCounter() + 1

	return slot
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
