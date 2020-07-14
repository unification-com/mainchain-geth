package dsg

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/consensus"
	"github.com/unification-com/mainchain/log"
	"github.com/unification-com/mainchain/p2p"
)

func (d *Dsg) SetProtocolManager(pm consensus.DsgProtocolManager) {
	log.Trace("dsg SetProtocolManager")

	d.coreLock.Lock()
	defer d.coreLock.Unlock()

	d.pm = pm
}

func (d *Dsg) HandleMsg(addr common.Address, msg p2p.Msg) (bool, error) {
	d.coreLock.Lock()
	defer d.coreLock.Unlock()

	log.Info("dsg.HandleMsg - incoming msg", "msg.code", msg.Code, "msg.size", msg.Size)

	if !d.engineRunning {
		return true, errEngineNotRunning
	}

	numInvalids := d.dsgCache.GetInvalidCounter()
	slot := d.chain.CurrentHeader().SlotCount + 1 + numInvalids
	isEv := d.Member(slot, d.signer)

	switch {
	case msg.Code == ValidationMsg:
		if !isEv {
			log.Info("dsg.HandleMsg - message not for me - ignore")
			return true, nil
		}
		// Validation Message indicates whether the peers agree with the BP or not
		var validationMessage ValidationMessage
		if err := msg.Decode(&validationMessage); err != nil {
			return true, err
		}

		log.Info("Validation Result:", "block", validationMessage.Number, "proposer", validationMessage.Proposer, "validator", validationMessage.Verifier, "authorise", validationMessage.Authorize)

		if !d.Member(slot, validationMessage.Verifier) {
			log.Info("discarding validation message from invalid EV", "slot", slot, "ev", validationMessage.Verifier)
			return true, nil
		}

		result := d.dsgCache.InsertValidationMessage(validationMessage)

		evid := d.EVIdFromEtherbase(d.signer)

		// Todo - add validation signature to extraData

		if result == Accept {
			// Post result event to handle timers
			go d.dsgEventMux.Post(ValidationResultEvent{Valid: true})

            // get block from cache
			blockProposal, err := d.dsgCache.GetBlockProposal(validationMessage.Number, validationMessage.Proposer)

			// if the accepted block was this EV's, post the BlockVerifiedEvent to send
			// final block back to engine.Seal
			if evid == d.EVIdFromEtherbase(validationMessage.Proposer) {
				log.Info("Network accepting my block", "Block Number", validationMessage.Number, "Proposer", validationMessage.Proposer, "txs", len(blockProposal.ProposedBlock.Transactions()))
				if err != nil {
					log.Error("Block Proposal not found in cache", "err", err)
					return true, err
				}
				go d.dsgEventMux.Post(BlockVerifiedEvent{BlockProposal: &blockProposal})
			} else {
				// todo - need some way of telling the EV who's block was accepted to post back to engine.Seal
			}
		}
		if result == Reject {
			go d.dsgEventMux.Post(ValidationResultEvent{Valid: false})
			d.dsgCache.IncrementInvalidCounter()
			rbp := RequestNewBlockProposalMessage{
				Number:    validationMessage.Number,
				Verifier:  d.signer,
				Signature: common.Hash{},
			}
			log.Info("request new block proposal - reject invalid BP", "num", rbp.Number, "proposer", validationMessage.Proposer)
			go d.dsgEventMux.Post(RequestNewBlockProposalEvent{ RequestNewBlockProposalMessage: &rbp})
			go d.dsgEventMux.Post(RequestNewBlockProposalMessageEvent{RequestNewBlockProposalMessage: &rbp})
		}
		return true, nil

	case msg.Code == BlockProposalMsg:
		// A new block has been proposed
		if !isEv {
			log.Info("dsg.HandleMsg - message not for me - ignore")
			return true, nil
		}

		var proposal BlockProposal
		if err := msg.Decode(&proposal); err != nil {
			return true, err
		}

		if !d.Member(slot, proposal.Proposer) {
			log.Info("discarding block proposal message from invalid EV", "slot", slot, "ev", proposal.Proposer)
			return true, nil
		}

		// cache BP
		d.dsgCache.InsertBlockProposal(proposal)
		valid := d.ValidateBlockProposal(proposal)
		log.Info("BlockProposal Received", "num", proposal.Number, "proposer", proposal.Proposer, "txs", len(proposal.ProposedBlock.Transactions()), "valid", valid)

		go d.dsgEventMux.Post(NewBlockProposalFoundEvent{valid})

		vm := ValidationMessage{Number: proposal.Number, BlockHash: proposal.BlockHash, Verifier:d.signer, Proposer: proposal.Proposer, Signature: common.Hash{}, Authorize:valid}

		// cache own validation message
		d.dsgCache.InsertValidationMessage(vm)
		go d.dsgEventMux.Post(SendNewValidationMessageEvent{ValidationMessage: &vm})

		return true, nil

	case msg.Code == RequestNewBlockProposalMsg:
		if !isEv {
			log.Info("dsg.HandleMsg - message not for me - ignore")
			return true, nil
		}
		var requestProposal RequestNewBlockProposalMessage
		if err := msg.Decode(&requestProposal); err != nil {
			return true, err
		}

		if !d.Member(slot, requestProposal.Verifier) {
			log.Info("discarding request new block proposal message from invalid EV", "slot", slot, "ev", requestProposal.Verifier)
			return true, nil
		}
		log.Info("A new block proposal has been requested", "num", requestProposal.Number, "from", requestProposal.Verifier)

		go d.dsgEventMux.Post(RequestNewBlockProposalEvent{ RequestNewBlockProposalMessage: &requestProposal})

		return true, nil
	default:
		log.Info("dsg.HandleMsg - not a DSG message. Pass back to eth/handler")
		// Not a DSG message. Pass back to eth/handler
		return false, nil
	}

	// Not a DSG message. Pass back to eth/handler
	return false, nil
}

func (d *Dsg) NewChainHead() error {
	log.Info("dsgHandler - NewChainHead")
	return nil
}

func (d *Dsg) broadcastLoop() {
	for {
		select {
		case obj := <-d.sendProposalBlockSub.Chan():
			if ev, ok := obj.Data.(SendNewBlockProposalEvent); ok {
				log.Info("broadcast SendNewBlockProposalEvent")
				inTurnPeers := d.getInRoundEvs(1)
				if d.pm != nil && len(inTurnPeers) > 0 {
					peers := d.pm.FindPeers(inTurnPeers)
					for _, p := range peers {
						go p.Send(BlockProposalMsg, ev.BlockProposal)
					}
				}
			}
		case obj := <-d.sendRequestNewBlockProposalSub.Chan():
			if ev, ok := obj.Data.(RequestNewBlockProposalMessageEvent); ok {
				log.Info("broadcast RequestNewBlockProposalMessageEvent")
				inTurnPeers := d.getInRoundEvs(1)
				if d.pm != nil && len(inTurnPeers) > 0 {
					peers := d.pm.FindPeers(inTurnPeers)
					for _, p := range peers {
						go p.Send(RequestNewBlockProposalMsg, ev.RequestNewBlockProposalMessage)
					}
				}
			}
		case obj := <-d.sendNewValidationMessageSub.Chan():
			if ev, ok := obj.Data.(SendNewValidationMessageEvent); ok {
				log.Info("broadcast SendNewValidationMessageEvent")
				inTurnPeers := d.getInRoundEvs(1)
				if d.pm != nil && len(inTurnPeers) > 0 {
					peers := d.pm.FindPeers(inTurnPeers)
					for _, p := range peers {
						go p.Send(ValidationMsg, ev.ValidationMessage)
					}
				}
			}
		}
	}
}
