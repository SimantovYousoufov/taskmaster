package data

import "github.com/pkg/errors"

const (
	MITLimit = 3
	TodoLimit = 10
)
var (
	ErrAtTaskLimit = errors.New("cannot add any more tasks to list")
	ErrOutOfBounds = errors.New("out of bounds")
)
