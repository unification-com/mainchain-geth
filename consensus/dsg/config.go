package dsg

type Config struct {
	ValidatorPool      uint64         `json:"pool"` // The number of initial validators in the epoch's pool
	ActiveSigners      uint64         `json:"signers"` // The number of EVs per round
	BlocksInEpoch      uint64         `json:"epoch"` // The number of blocks in an epoch
	EpochSubdivisions  uint64         `json:"division"` // The number of rounds per epoch
}

// String implements the stringer interface, returning the consensus engine details.
func (c *Config) String() string {
	return "dsg"
}
