package pool

import (
	"sync"
	. "github.com/pool-beta/pool-server/types"
)


// PoolFactory contains the API for managing Pools (singleton)
type PoolFactory interface {
	// Creates a new Pool
	CreateNewPool(owner UserID, poolName string) (Pool, error)
	RetreivePool(PoolID) (Pool, error)
}

type poolFactory struct {
	pools map[PoolID]*pool
	mutex sync.Mutex
}

func NewPoolFactory() (PoolFactory, error) {
	pools := make(map[PoolID]*pool)

	return &poolFactory{
		pools: pools,
	}, nil
}

func (pFactory *poolFactory) CreateNewPool(owner UserID, poolName string) (Pool, error) {
	// TODO: Validate owner

	// Create a new pool id
	pid := NewPoolID()

	pool := initPool(pid, poolName, owner)
	return pool, nil
}

func (pFactory *poolFactory) RetreivePool(pid PoolID) (Pool, error) {
	return nil, nil
}


//--------------------------------------------------------------------------------------------------

type Pool interface {
	// Money Power
	AddPusher(UserID) error
	RemovePusher(UserID) error
	AddPuller(UserID) error
	RemovePuller(UserID) error

	// Admin Power
	AddOwner(UserID) error
	RemoveOwner(UserID) error
	AddAdmin(UserID) error
	RemoveAdmin(UserID) error
	AddMember(UserID) error
	RemoveMember(UserID) error

	// Triggers (?)
	Pull() error
	Push() error
	Reset() error // (?)
}

type pool struct {
	id PoolID
	name string

	pushers []UserID
	pullers []UserID
	owners []UserID
	admins []UserID
	members []UserID

	mutex sync.Mutex
}

// Used by PoolFactory to init a Pool; error check should be already done
func initPool(id PoolID, name string, owner UserID) *pool {
	// Init
	pushers := make([]UserID, 0)
	pullers := make([]UserID, 0)
	owners := make([]UserID, 1)
	admins := make([]UserID, 0)
	members := make([]UserID, 0)

	// Add Original Owner
	owners[0] = owner

	return &pool {
		id: id,
		name: name,

		pushers: pushers,
		pullers: pullers,
		owners: owners,
		admins: admins,
		members: members,
	}
}

func (p *pool) AddPusher(pusher UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Validate Pusher

	p.pushers = append(p.pushers, pusher) 
	return nil
}

func (p *pool) RemovePusher(pusher UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Validate Pusher

	p.pushers = append(p.pushers, pusher) 

	return nil
}

func (p *pool) AddPuller(UserID) error {
	return nil
}

func (p *pool) RemovePuller(UserID) error {
	return nil
}

func (p *pool) AddOwner(UserID) error {
	return nil
}

func (p *pool) RemoveOwner(UserID) error {
	return nil
}

func (p *pool) AddAdmin(UserID) error {
	return nil
}

func (p *pool) RemoveAdmin(UserID) error {
	return nil
}

func (p *pool) AddMember(UserID) error {
	return nil
}

func (p *pool) RemoveMember(UserID) error {
	return nil
}

func (p *pool) Push() error {
	return nil
}

func (p *pool) Pull() error {
	return nil
}

func (p *pool) Reset() error {
	return nil
}


