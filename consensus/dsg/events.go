package dsg

// NewBlockValidatedEvent: used to manager timeouts
type NewBlockValidatedEvent struct{ }

// NewBlockProposalFoundEvent: used to manage timeouts
type NewBlockProposalFoundEvent struct{ Valid bool }

// SendNewBlockProposalEvent: used to instruct the ProtocolManager to send a BlockProposal
type SendNewBlockProposalEvent struct{ BlockProposal *BlockProposal }

// RequestNewBlockProposalMessageEvent
type RequestNewBlockProposalMessageEvent struct{ RequestNewBlockProposalMessage *RequestNewBlockProposalMessage }

// BlockVerifiedEvent: consensus found on block and trigger NewBlockMessage
type BlockVerifiedEvent struct{ BlockProposal *BlockProposal }

// ValidationResultEvent: used to manage timeouts for TS5 and TS6
type ValidationResultEvent struct{ Valid bool }

// RequestNewBlockProposalEvent - used to check for and trigger a new BP
type RequestNewBlockProposalEvent struct { RequestNewBlockProposalMessage *RequestNewBlockProposalMessage }

//SendNewValidationMessageEvent - used to broadcast Validation Messages to peers
type SendNewValidationMessageEvent struct { ValidationMessage *ValidationMessage }
