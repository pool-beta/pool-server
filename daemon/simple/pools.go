package simple

import (
	ppool "github.com/pool-beta/pool-server/pool"
)

/*
	Implements POOL (simple pools)
*/

type pools struct {
	pf ppool.PoolFactory
}

type pool struct {
	pool ppool.Pool
}

type drain struct {
	drain ppool.Drain
}

func NewPools() (Pools, error) {
	pf, err := ppool.NewPoolFactory()
	if err != nil {
		return nil, err
	}

	return &pools{
		pf: pf,
	}, nil
}

func (ps *pools) NewPool() Pool {
	p := ps.pf.CreatePool()

	return &pool{
		pool: p,
	}
}

func (ps *pools) NewDrainPool() Drain {
	d := ps.pf.CreateDrain()

	return &drain{

	}
}
	NewTankPool() Tank
	RemovePool() error

	NewStream() Stream
	RemoveStream() error