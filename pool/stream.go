package pool

import (
	"fmt"

	. "github.com/pool-beta/pool-server/types"
)

type Stream interface {
	Pull(Drop) error
	Push(Drop) error
}

type stream struct {
	owner UserID
	// Source
	pullPool Pool
	// Destination
	pushPool Pool
}

func NewStream(owner UserID, pullPool Pool, pushPool Pool) (Stream, error) {
	if pushPool == nil || pullPool == nil {
		return nil, fmt.Errorf("Invalid pushPool or pullPool")
	}

	return &stream{
		owner: owner,
		pullPool: pullPool,
		pushPool: pushPool,
	}, nil
}

func (s *stream) Pull(drop Drop) error {
	return s.pullPool.Pull(drop)
}

func (s *stream) Push(drop Drop) error {
	return nil
}
