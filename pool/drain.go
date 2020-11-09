package pool

import (
	"fmt"

	. "github.com/pool-beta/pool-server/types"
)

type Drain interface {
	// Extends Pool
	Pool
}

type drain struct {
	*pool
}

func NewDrain(name string, owner UserID) Drain {
	pid := NewPoolID()

	p := initPool(pid, name, owner)
	p.reserve = USDollar(0)

	return &drain{
		pool: p,
	}
}

func (d *drain) Push(drop Drop) error {
	return fmt.Errorf("Can't push from a drain -- drop: %v", drop)
}