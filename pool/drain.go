package pool

import (
	"fmt"

	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/user/types"
)

/* 
	Drain is pool with no reserve, and cannot push
	It can only pull, and implements the debit card
*/

type Drain interface {
	// Extends Pool
	Pool
}

type drain struct {
	*pool
}

func newDrain(pid PoolID, name string, owner UserID) Drain {
	p := initPool(pid, name, owner)
	p.reserve = USDollar(0)

	return &drain{
		pool: p,
	}
}

func (d *drain) Push(drop Drop) error {
	return fmt.Errorf("Can't push from a drain -- drop: %v", drop)
}