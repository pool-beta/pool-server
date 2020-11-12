package pool

import (
	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/user/types"
)

type Tank interface {
	Pool
}

type tank struct {
	*pool
}

func newTank(pid PoolID, name string, owner UserID) Tank {
	p := initPool(pid, name, owner)
	p.reserve = USDollar(0) // never checked

	return &tank{
		pool: p,
	}
}

func (t *tank) PullDrop(drop Drop, useReserve bool) error {
	// Should allow infinite pulls

	// TODO: log the pull
	return nil
}