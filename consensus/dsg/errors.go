package dsg

import "errors"

var (
	errGetBPFromCacheFailed = errors.New("could not retrieve block proposal from cache")
	errCannotSealGenesis    = errors.New("cannot seal genesis block")
	errNotAuthorisedToPropose = errors.New("ev not authorised to propose")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte signature suffix missing")
)
