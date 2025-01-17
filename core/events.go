// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/consensus/dsg"
	"github.com/unification-com/mainchain/core/types"
)

// NewTxsEvent is posted when a batch of transactions enter the transaction pool.
type NewTxsEvent struct{ Txs []*types.Transaction }

// PendingLogsEvent is posted pre mining and notifies of pending logs.
type PendingLogsEvent struct {
	Logs []*types.Log
}

// NewMinedBlockEvent is posted when a block has been imported.
type NewMinedBlockEvent struct{ Block *types.Block }

// NewBlockValidatedEvent: used to manager timeouts
type NewBlockValidatedEvent struct{ }

// NewBlockProposalFoundEvent: used to manage timeouts
type NewBlockProposalFoundEvent struct{ Valid bool }

// NewBlockProposalEvent: used to instruct the ProtocolManager to send a BlockProposal
type NewBlockProposalEvent struct{ BlockProposal *dsg.BlockProposal }

// RequestNewBlockProposalMessageEvent
type RequestNewBlockProposalMessage struct{ RequestNewBlockProposalMessage *dsg.RequestNewBlockProposalMessage }

// BlockVerifiedEvent: consensus found on block and trigger NewBlockMessage
type BlockVerifiedEvent struct{ BlockProposal *dsg.BlockProposal }

// ValidationResultEvent: used to manage timeouts for TS5 and TS6
type ValidationResultEvent struct{ Valid bool }

// RequestNewBlockProposalEvent - used to check for and trigger a new BP
type RequestNewBlockProposalEvent struct { RequestNewBlockProposalMessage *dsg.RequestNewBlockProposalMessage }

//SendNewValidationMessageEvent - used to broadcast Validation Messages to peers
type SendNewValidationMessageEvent struct { ValidationMessage *dsg.ValidationMessage }

// RemovedLogsEvent is posted when a reorg happens
type RemovedLogsEvent struct{ Logs []*types.Log }

type ChainEvent struct {
	Block *types.Block
	Hash  common.Hash
	Logs  []*types.Log
}

type ChainSideEvent struct {
	Block *types.Block
}

type ChainHeadEvent struct{ Block *types.Block }
