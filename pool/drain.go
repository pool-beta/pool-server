package pool

import (
	"fmt"

	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/user/types"
)

/*
	Drain is pool with no reserve, and cannot push
	It implements the POOL debit card
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

func (d *drain) Push(flow Flow) error {
	if flow != nil {
		flow.Invalid()
	}

	return fmt.Errorf("A drain cannot push")
}

func (d *drain) Pull(flow Flow) error {
	err := d.PullDrop(flow.PullDrop(), false)
	if err != nil {
		flow.Invalid()
		return err
	}

	// Initiate Push if necessary
	err = d.PushDrop(flow.PushDrop(), true)
	if err != nil {
		flow.Invalid()
		return err
	}

	flow.Valid()
	return nil
}

func (d *drain) PushDrop(drop Drop, useReserve bool) error {
	// drain does not keep a reserve
	// Accept all pushes
	drop.AddWithheld(USDollar(0))
	return nil
}

func (d *drain) GetType() string {
	return "drain"
}
