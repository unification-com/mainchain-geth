package dsg

import "errors"

var (
	errGetBPFromCacheFailed = errors.New("could not retrieve block proposal from cache")
)
