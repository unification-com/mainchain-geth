package dsg

import (
	"bytes"
	"github.com/unification-com/mainchain/accounts"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/consensus"
	"github.com/unification-com/mainchain/core/state"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/event"
	"github.com/unification-com/mainchain/log"
	"github.com/unification-com/mainchain/params"
	"github.com/unification-com/mainchain/rlp"
	"github.com/unification-com/mainchain/rpc"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
	"sync"
	"time"
)

var (
	defaultBlockTime = 5 // default block time if none is set in DSG Config
	dsgExtraVanity   = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	dsgExtraSeal     = 65 // Fixed number of extra-data suffix bytes reserved for Sealing EV's signature

	blockProposalTimeoutDuration = 3 * time.Second
	validationTimeoutDuration    = 3 * time.Second

)

// SignerFn is a signer callback function to request a header to be signed by a
// backing account.
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

type Dsg struct {
	config       *params.DsgConfig     // Consensus engine configuration parameters
	dsgCache     *Cache                // DSG Cache
	chain        consensus.ChainReader
	currentBlock func() *types.Block

	pm consensus.DsgProtocolManager

	signer   common.Address // Ethereum address of the signing key
	signFn   SignerFn       // Signer function to authorize hashes with
	lock     sync.RWMutex   // Protects the signer fields
	coreLock sync.RWMutex   // used during engine start/stop

	dsgEventMux *event.TypeMux

	// Channels
	verifiedBlockCh chan *VerifiedBlock

	// Receiver subscriptions
	receiveVerifiedBlockSub           *event.TypeMuxSubscription
	receiveNewBlockValidSub           *event.TypeMuxSubscription
	receiveNewBlockProposalSub        *event.TypeMuxSubscription
	receiveValidationResultSub        *event.TypeMuxSubscription
	receiveRequestNewBlockProposalSub *event.TypeMuxSubscription

	// Broadcaster subscriptions
	sendRequestNewBlockProposalSub *event.TypeMuxSubscription
	sendProposalBlockSub           *event.TypeMuxSubscription
	sendNewValidationMessageSub    *event.TypeMuxSubscription

	// Timers
	timeoutBlockProposal *time.Timer
	validationTimeout    *time.Timer

	engineRunning bool
}

// NewEngine returns a new DSG consensus.Engine
func NewEngine(config *params.DsgConfig) consensus.DsgEngine {

	conf := *config
	if conf.BlockTime == 0 {
		conf.BlockTime = uint64(defaultBlockTime)
	}

	dsg := &Dsg{
		config:          &conf,
		dsgCache:        NewCache(),
		dsgEventMux:     new(event.TypeMux),
		verifiedBlockCh: make(chan *VerifiedBlock),
		engineRunning:   false,
	}

	// Receiver subscriptions
	dsg.receiveVerifiedBlockSub = dsg.dsgEventMux.Subscribe(BlockVerifiedEvent{})
	dsg.receiveNewBlockValidSub = dsg.dsgEventMux.Subscribe(NewBlockValidatedEvent{})
	dsg.receiveNewBlockProposalSub = dsg.dsgEventMux.Subscribe(NewBlockProposalFoundEvent{})
	dsg.receiveValidationResultSub = dsg.dsgEventMux.Subscribe(ValidationResultEvent{})
	dsg.receiveRequestNewBlockProposalSub = dsg.dsgEventMux.Subscribe(RequestNewBlockProposalEvent{})

	// Broadcaster subscriptions
	dsg.sendRequestNewBlockProposalSub = dsg.dsgEventMux.Subscribe(RequestNewBlockProposalEvent{})
	dsg.sendProposalBlockSub = dsg.dsgEventMux.Subscribe(SendNewBlockProposalEvent{})
	dsg.sendNewValidationMessageSub =dsg.dsgEventMux.Subscribe(SendNewValidationMessageEvent{})

	// Timers
	dsg.timeoutBlockProposal = time.NewTimer(blockProposalTimeoutDuration)
	dsg.validationTimeout = time.NewTimer(validationTimeoutDuration)

	return dsg
}

func (d *Dsg) Start(chain consensus.ChainReader, currentBlock func() *types.Block) error {
	log.Trace("start DsgEngine")
	d.coreLock.Lock()
	d.chain = chain
	d.currentBlock = currentBlock

	if !d.engineRunning {
		d.engineRunning = true
	}
	d.coreLock.Unlock()

	go d.coreLoop()
	go d.broadcastLoop()

	return nil
}

func (d *Dsg) Stop() error {
	log.Trace("stop DsgEngine")
	d.coreLock.Lock()
	defer d.coreLock.Unlock()
	d.engineRunning = false

	return nil
}

// Author implements engine.Author
func (d *Dsg) Author(header *types.Header) (common.Address, error) {
	log.Info("engine.Author", "header.Number", header.Number)
	// Todo - get Author from DsgExtra struct{}
	return ecrecover(header)
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

	go func() {
		for _, _ = range headers {
			select {
			case <-abort:
				return
			case results <- nil:
			}
		}
	}()

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
// header is the block to be proposed
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
	// Todo - change to RLP encoded empty DsgExtra struct{} containing EVs, seals etc.
	if len(header.Extra) < dsgExtraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, dsgExtraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:dsgExtraVanity]
	// append 65 bytes
	header.Extra = append(header.Extra, make([]byte, dsgExtraSeal)...)

	// Set block's expected timestamp
	header.Time = parent.Time + d.config.BlockTime
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	// Set block's expected slot number
	// new block, so shouldn't have any invalids yet - just parent slot + 1
	header.SlotCount = parent.SlotCount + 1

	// Todo - look into perhaps resetting slot counter to 1 for each new epoch

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
	log.Info("engine.FinalizeAndAssemble", "header.Number", header.Number, "txs", len(txs))

	// Todo - add block rewards to EV before setting header.Root - will need header.Coinbase for this
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Uncles are not used in DSG
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// Seal implements engine.Seal
// Receives the new block for sealing from the worker, then passes it to the DSG backend
// for proposing & validation. Result is returned back to the worker on results chan<-
func (d *Dsg) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	log.Info("engine.Seal", "block.Header()", block.Header().Number, "txs", len(block.Transactions()))

	header := block.Header()

	// No point sealing genesis
	if header.Number.Uint64() == 0 {
		return errCannotSealGenesis
	}

	d.lock.RLock()
	signer, signFn := d.signer, d.signFn
	d.lock.RUnlock()

	slot := header.SlotCount // set during engine.Prepare. Should just be parent.SlotCount + 1

	// return error if we're not authorised to propose & seal - i.e. not a member fo the current round
	member := d.Member(slot, d.signer)
	if !member {
		return errNotAuthorisedToPropose
	}

	// check if it's our turn to seal & propose
	v := d.Authorized(header)

	// set delay to specified block time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now())

	// Sign
	// Todo - modify extraData to DsgExtra struct{} containing EVs, seals etc.
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeClique, DSGRLP(header)) // Todo - MimetypeClique?
	if err != nil {
		return err
	}

	// inject the signature
	copy(header.Extra[len(header.Extra)-dsgExtraSeal:], sighash)

	// seal
	block = block.WithSeal(header)

	// generate proposal & cache locally
	blockProposal := d.ProposeBlock(block, signer)
	d.dsgCache.InsertBlockProposal(blockProposal)

	// no need to run until block proposal process has begun
	d.timeoutBlockProposal.Stop()
	d.validationTimeout.Stop()

	log.Info("delay until block time", "delay",  common.PrettyDuration(delay))

	// tmp delay
	tmpDelay := time.NewTimer(15 * time.Second) // todo: get rid of this egregious hack
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
			if v {
				log.Info("The EV is in turn to be first to propose this block")
				//validate, cache and broadcast own validation message
				valid := d.ValidateBlockProposal(blockProposal)
				vm := ValidationMessage{Number: blockProposal.Number, BlockHash: blockProposal.BlockHash, Verifier: signer, Proposer: blockProposal.Proposer, Signature: common.Hash{}, Authorize:valid}
				d.dsgCache.InsertValidationMessage(vm)
				go d.dsgEventMux.Post(SendNewValidationMessageEvent{ValidationMessage:&vm})

				// broadcast BP, and check for incoming validation messages
				go d.dsgEventMux.Post(SendNewBlockProposalEvent{BlockProposal: &blockProposal})
				go d.dsgEventMux.Post(NewBlockProposalFoundEvent{ Valid: valid})
			} else {
				// Not this EV's turn, start timer and wait for incoming BP
				d.timeoutBlockProposal.Reset(blockProposalTimeoutDuration)
			}
		}

		select {
		case verifiedBlock := <-d.verifiedBlockCh:
			// block verified be EVs - return it to worker's results chan<-
			log.Info("engine.Seal - verifiedBlock received. pass block back to worker", "num", verifiedBlock.FinalBlock.Number(), "txs", len(verifiedBlock.FinalBlock.Transactions()))
		    results <- verifiedBlock.FinalBlock.WithSeal(verifiedBlock.FinalBlock.Header())
		    return
		case <-tmpDelay.C: // todo: get rid of this egregious hack
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()
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

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (d *Dsg) Authorize(signer common.Address, signFn SignerFn) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.signer = signer
	d.signFn = signFn
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
		header.Extra[:len(header.Extra)-dsgExtraSeal], // Yes, this will panic if extra is too short
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

func ecrecover(header *types.Header) (common.Address, error) {
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < dsgExtraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-dsgExtraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	return signer, nil
}
