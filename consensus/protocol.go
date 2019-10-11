package consensus

import "github.com/unification-com/mainchain/common"

type DsgProtocolManager interface {
	FindPeers(map[common.Address]bool) map[common.Address]Peer
}

// Peer defines the interface to communicate with peer
type Peer interface {
	// Send sends the message to this peer
	Send(msgcode uint64, data interface{}) error
}