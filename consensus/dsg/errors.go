package dsg

import "errors"

var (
	errGetBPFromCacheFailed = errors.New("could not retrieve block proposal from cache")
	errCannotSealGenesis    = errors.New("cannot seal genesis block")
	errNotAuthorisedToPropose = errors.New("ev not authorised to propose")
)
