package dsg

type Config struct {
	NumSignersInEpoch uint64 `json:"epochsize"`  // The number of initial validators in the epoch's pool
	NumSignersinRound uint64 `json:"roundsize"`  // The number of EVs per round
	BlocksInEpoch     uint64 `json:"epoch"`      // The number of blocks in an epoch
	NumberOfRounds    uint64 `json:"round"`      // The number of rounds per epoch
}

// String implements the stringer interface, returning the consensus engine details.
func (c *Config) String() string {
	return "dsg"
}
