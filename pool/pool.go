package pool

import (
	"fmt"
	"sync"

	. "github.com/pool-beta/pool-server/types"
)


// PoolFactory contains the API for managing Pools (singleton; manually managed)
type PoolFactory interface {
	// Creates a new Pool
	CreatePool(string, UserID) (Pool, error)
	RetrievePool(PoolID) (Pool, error)
	ReturnPool(PoolID) error
}

type poolFactory struct {
	pools map[PoolID]*poolRef
	mutex sync.Mutex
}

// Interal struct for PoolFactory to keep track of active Pools
// TODO: Flushing mechanism of some sort
type poolRef struct {
	pool Pool
	refCount int
	mutex sync.Mutex
}

func NewPoolFactory() (PoolFactory, error) {
	pools := make(map[PoolID]*poolRef)

	return &poolFactory{
		pools: pools,
	}, nil
}

func (pf *poolFactory) CreatePool(poolName string, owner UserID) (Pool, error) {
	// TODO: Validate owner

	// Create a new pool id
	pid := NewPoolID()

	pool := initPool(pid, poolName, owner)

	// Create poolRef
	pr := &poolRef{
		pool: pool,
		refCount: 1,
	}

	pf.pools[pid] = pr
	return pool, nil
}

func (pf *poolFactory) RetrievePool(pid PoolID) (Pool, error) {
	pr, ok := pf.pools[pid]
	if !ok {
		return nil, fmt.Errorf("Pool does not exist -- pool: %v", pid)
	}

	pr.mutex.Lock()
	defer pr.mutex.Unlock()

	pr.refCount++
	return pr.pool, nil
}

func (pf *poolFactory) ReturnPool(pid PoolID) error {
	pr, ok := pf.pools[pid]
	if !ok {
		return fmt.Errorf("Pool does not exist -- pool: %v", pid)
	}

	pr.mutex.Lock()
	defer pr.mutex.Unlock()

	pr.refCount--

	// TODO: Flush when refCount == 0
	return nil
}


//--------------------------------------------------------------------------------------------------

type Pool interface {
	// Auth
	PoolAuth

	// Getters
	GetReserve() USDollar

	// Drop Control
	AddPusher(Stream) error
	RemovePusher(Stream) error
	AddPuller(Stream) error
	RemovePuller(Stream) error

	// Drop
	Pull(Drop) error
	Push(Drop) error
	Reset() error // (?)

	// For Tests
	Fund(USDollar) USDollar
}

type pool struct {
	*poolAuth
	id PoolID
	name string
	reserve USDollar

	pushers []Stream
	pullers []Stream

	mutex sync.Mutex
}

// Used by PoolFactory to init a Pool; error check should be already done
func initPool(id PoolID, name string, owner UserID) *pool {
	// Initialize
	pushers := make([]Stream, 0)
	pullers := make([]Stream, 0)

	// Creat PoolAuth
	auth := NewPoolAuth(id)

	// Add Original Owner
	auth.AddOwner(owner)

	return &pool {
		id: id,
		name: name,
		reserve: USDollar(0), // Always start with zero

		pushers: pushers,
		pullers: pullers,

		poolAuth: auth,
	}
}

// Triggers
func (p *pool) Push(drop Drop) error {
	return nil
}

func (p *pool) Pull(drop Drop) error {
	// Pull from pullers if needed

	return nil
}

func (p *pool) Reset() error {
	return nil
}

func (p *pool) AddPusher(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.pushers = append(p.pushers, stream) 
	return nil
}

func (p *pool) RemovePusher(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	pushers, ok := FindAndRemove(p.pushers, stream)
	if !ok {
		return fmt.Errorf("User is not a pusher of this pool -- userID: %v; poolID: %v", stream, p.id)
	}
	p.pushers = pushers
	return nil
}

func (p *pool) AddPuller(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.pullers = append(p.pullers, stream) 
	return nil
}

func (p *pool) RemovePuller(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	pullers, ok := FindAndRemove(p.pullers, stream)
	if !ok {
		return fmt.Errorf("User is not a puller of this pool -- userID: %v; poolID: %v", stream, p.id)
	}
	p.pullers = pullers
	return nil
}

// Getters
func (p *pool) GetReserve() USDollar {
	return p.Fund(USDollar(0))
}

// For tesing
func (p *pool) Fund(amount USDollar) USDollar {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.reserve += amount
	return p.reserve
}




