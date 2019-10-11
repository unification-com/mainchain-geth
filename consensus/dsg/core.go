package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/log"
	"math/big"
)

const (
    // Protocol messages belonging to DSG
    BlockProposalMsg           = 0x11
    ValidationMsg              = 0x12
    RequestNewBlockProposalMsg = 0x13
)

func (d *Dsg) coreLoop() {
	defer d.receiveVerifiedBlockSub.Unsubscribe()
	defer d.receiveNewBlockValidSub.Unsubscribe()
	defer d.receiveNewBlockProposalSub.Unsubscribe()
	d.validationTimeout.Stop()

	for {
		select {
		case obj := <-d.receiveVerifiedBlockSub.Chan():
			if obj != nil {
				if ev, ok := obj.Data.(BlockVerifiedEvent); ok {
					// network accepted this EV's block. Send back to engine.Seal
					select {
					case d.verifiedBlockCh <- &VerifiedBlock{FinalBlock: ev.BlockProposal.ProposedBlock}:
					}
				}
			}

		// Todo - this never seems to be reached
		case obj := <-d.receiveNewBlockValidSub.Chan():
			if obj != nil {
				if ev, ok := obj.Data.(NewBlockValidatedEvent); ok {
					log.Info("NewBlockValidatedEvent", "ev", ev)
					d.timeoutBlockProposal.Reset(blockProposalTimeoutDuration)
					d.dsgCache.ResetInvalidCounter()
				}
			}

		case obj := <-d.receiveNewBlockProposalSub.Chan():
			if obj != nil {
				if ev, ok := obj.Data.(NewBlockProposalFoundEvent); ok {
					log.Info("NewBlockProposalFoundEvent", "valid", ev.Valid)
					if ev.Valid {
						d.validationTimeout.Reset(validationTimeoutDuration)
					}
					d.timeoutBlockProposal.Stop()
				}
			}

		case obj := <-d.receiveValidationResultSub.Chan():
			if obj != nil {
				if ev, ok := obj.Data.(ValidationResultEvent); ok {
					log.Info("ValidationResultEvent", "ev", ev)
					if ev.Valid == true {
						d.timeoutBlockProposal.Stop()
						d.validationTimeout.Stop()
					}
					if ev.Valid == false {
						d.timeoutBlockProposal.Reset(blockProposalTimeoutDuration)
						d.validationTimeout.Stop()
					}
				}
			}

		case obj := <-d.receiveRequestNewBlockProposalSub.Chan():
			if obj != nil {
				if ev, ok := obj.Data.(RequestNewBlockProposalEvent); ok {
					bpRequired := d.dsgCache.InsertRequestNewBlockProposalMessage(*ev.RequestNewBlockProposalMessage)

					// Get my BP from cache
					bp, err := d.dsgCache.GetBlockProposal(ev.RequestNewBlockProposalMessage.Number, d.signer)
					if err != nil {
						log.Info(" error - could not find my bp in cache", "num", ev.RequestNewBlockProposalMessage.Number, "err", err, "etherbase", d.signer)
					    return
					}

					v := d.Authorized(bp.ProposedBlock.Header())

					log.Info("receiveRequestNewBlockProposalSub.Chan", "bpRequired", bpRequired, "v", v)

					if bpRequired && v{
						log.Info("new block proposal request for me - I should broadcast my BP")
						// Todo - correct the newBlock stuff
						newBlock := d.SetSlotNumberInBlock(bp.ProposedBlock)
						blockProposal := d.ProposeBlock(newBlock, d.signer)

						//validate, cache and broadcast own validation message
						valid := d.ValidateBlockProposal(blockProposal)
						vm := ValidationMessage{Number: blockProposal.Number, BlockHash: blockProposal.BlockHash, Verifier: d.signer, Proposer: blockProposal.Proposer, Signature: common.Hash{}, Authorize:valid}
						d.dsgCache.InsertValidationMessage(vm)
						go d.dsgEventMux.Post(SendNewValidationMessageEvent{ValidationMessage:&vm})
						go d.dsgEventMux.Post(SendNewBlockProposalEvent{BlockProposal: &blockProposal})
						go d.dsgEventMux.Post(NewBlockProposalFoundEvent{ Valid: valid})
					}
				}
			}

		case timeoutBlockProposalFire := <-d.timeoutBlockProposal.C:
			log.Info("timeoutBlockProposal fired", "timer", timeoutBlockProposalFire)
			d.dsgCache.IncrementInvalidCounter()

			d.timeoutBlockProposal.Reset(blockProposalTimeoutDuration)

			requestBlockNumber := big.NewInt(1)
			requestBlockNumber = requestBlockNumber.Add(requestBlockNumber,  d.currentBlock().Number())
			numInvalids := d.dsgCache.GetInvalidCounter()
			slot := d.chain.CurrentHeader().SlotCount + 1 + numInvalids
			expectedSignerIndex, _ := d.EVSlot(slot)
			expectedSigner := d.EtherbaseFromEVId(expectedSignerIndex)

			rbp := RequestNewBlockProposalMessage{
				Number:    requestBlockNumber,
				Verifier:  d.signer,
				Signature: common.Hash{},
			}
			log.Info("request new block proposal", "num", rbp.Number, "numInv", numInvalids, "slot", slot, "expectedSignerIndex", expectedSignerIndex, "expectedSigner", expectedSigner)

			// post RequestNewBlockProposalEvent and broadcast RequestNewBlockProposalMessageEvent
			go d.dsgEventMux.Post(RequestNewBlockProposalEvent{RequestNewBlockProposalMessage:&rbp})
			go d.dsgEventMux.Post(RequestNewBlockProposalMessageEvent{RequestNewBlockProposalMessage: &rbp})

		case validationTimeoutFire := <-d.validationTimeout.C:
			log.Info("validationTimeoutFire fired", "timer", validationTimeoutFire)

			requiredBlockNumber := big.NewInt(1)
			requiredBlockNumber = requiredBlockNumber.Add(requiredBlockNumber,  d.currentBlock().Number())

			numInvalids := d.dsgCache.GetInvalidCounter()
			slot := d.chain.CurrentHeader().SlotCount + 1 + numInvalids
			expectedSignerIndex, _ := d.EVSlot(slot)
			expectedSigner := d.EtherbaseFromEVId(expectedSignerIndex)
			status := d.dsgCache.QueryValidationState(requiredBlockNumber, expectedSigner)

			if status == Unknown || status == Reject {
				d.dsgCache.IncrementInvalidCounter()
				d.timeoutBlockProposal.Reset(blockProposalTimeoutDuration)
				d.validationTimeout.Stop()

				rbp := RequestNewBlockProposalMessage{
					Number:    requiredBlockNumber,
					Verifier:  d.signer,
					Signature: common.Hash{},
				}
				log.Info("request new block proposal", "num", rbp.Number, "numInv", numInvalids, "slot", slot, "expectedSignerIndex", expectedSignerIndex, "expectedSigner", expectedSigner)

				go d.dsgEventMux.Post(RequestNewBlockProposalEvent{RequestNewBlockProposalMessage:&rbp})
				go d.dsgEventMux.Post(RequestNewBlockProposalMessageEvent{RequestNewBlockProposalMessage: &rbp})
			}
			if status == Accept {
				d.timeoutBlockProposal.Stop()
				d.validationTimeout.Stop()
			}
		}
	}
}
