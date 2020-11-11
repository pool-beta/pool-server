package simple

import (
	ppool "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/pool/types"
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

type tank struct {
	tank ppool.Tank
}

func InitPools() (Pools, error) {
	pf, err := ppool.NewPoolFactory()
	if err != nil {
		return nil, err
	}

	return &pools{
		pf: pf,
	}, nil
}

func (ps *pools) CreatePool(user User, name string) (Pool, error) {
	p, err := ps.pf.CreatePool(name, user.ID(), POOL)
	if err != nil {
		return nil, err
	}
	
	return &pool{
		pool: p,
	}, nil
}

func (ps *pools) CreateDrainPool(user User, name string) (Drain, error) {
	d, err := ps.pf.CreatePool(name, user.ID(), DRAIN)
	if err != nil {
		return nil, err
	}

	return &drain{
		drain: d,
	}, nil
}

func (ps *pools) CreateTankPool(user User, name string) (Tank, error) {
	t, err := ps.pf.CreatePool(name, user.ID(), TANK)
	if err != nil {
		return nil, err
	}

	return &tank{
		tank: t,
	}, nil
}

func (ps *pools) RemovePool(pool Pool) error {
	
	return nil
}

func (ps *pools) CleanUp() error {
	return nil
}

// -------------------------------------------------------------------------------------------------

func (p *pool) CreateStream(puller Pool) (Stream, error) {
	return nil, nil
}

func (p *pool) RemoveStream(sid StreamID) error {
	return nil
}

func (p *pool) CreateFlow() (Flow, error) {
	return nil, nil
}

func (p *pool) CleanUp() error {
	return nil
}