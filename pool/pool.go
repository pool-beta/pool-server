package pool

import (
	"fmt"
	"sync"

	. "github.com/pool-beta/pool-server/types"
	"github.com/pool-beta/pool-server/utils"
)


// PoolFactory contains the API for managing Pools (singleton)
type PoolFactory interface {
	// Creates a new Pool
	CreatePool(owner UserID, poolName string) (Pool, error)
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

func (pf *poolFactory) CreatePool(owner UserID, poolName string) (Pool, error) {
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
	// Admin Control
	AdminCheck(user UserID, level string) bool
	AddOwner(UserID) error
	RemoveOwner(UserID) error
	AddAdmin(UserID) error
	RemoveAdmin(UserID) error
	AddMember(UserID) error
	RemoveMember(UserID) error

	// Drop Control
	AddPusher(Stream) error
	RemovePusher(Stream) error
	AddPuller(Stream) error
	RemovePuller(Stream) error

	// Drop
	Pull(Drop) error
	Push(Drop) error
	Reset() error // (?)
}

type pool struct {
	id PoolID
	name string
	reserve USDollar

	pushers []Stream
	pullers []Stream

	owners []UserID
	admins []UserID
	members []UserID

	mutex sync.Mutex
}

// Used by PoolFactory to init a Pool; error check should be already done
func initPool(id PoolID, name string, owner UserID) *pool {
	// Initialize
	pushers := make([]Stream, 0)
	pullers := make([]Stream, 0)

	owners := make([]UserID, 1)
	admins := make([]UserID, 0)
	members := make([]UserID, 0)

	// Add Original Owner
	owners[0] = owner

	return &pool {
		id: id,
		name: name,
		reserve: USDollar(0), // Always start with zero

		pushers: pushers,
		pullers: pullers,

		owners: owners,
		admins: admins,
		members: members,
	}
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


// Admin
func (p *pool) AdminCheck(user UserID, level string) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	
	ok := false
	switch level {
	case "owner":
		_, ok = utils.Find(p.owners, user)
	case "admin":
		_, ok = utils.Find(p.admins, user)
	case "member":
		_, ok = utils.Find(p.members, user)
	}
	return ok
}

func (p *pool) AddOwner(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.owners = append(p.owners, user) 
	return nil
}

func (p *pool) RemoveOwner(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	owners, ok := utils.FindAndRemove(p.owners, user)
	if !ok {
		return fmt.Errorf("User is not an owner of this pool -- userID: %v; poolID: %v", user, p.id)
	}
	p.owners = owners
	return nil
}

func (p *pool) AddAdmin(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.admins = append(p.admins, user) 
	return nil
}

func (p *pool) RemoveAdmin(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	admins, ok := utils.FindAndRemove(p.admins, user)
	if !ok {
		return fmt.Errorf("User is not an admin of this pool -- userID: %v; poolID: %v", user, p.id)
	}
	p.admins = admins
	return nil
}

func (p *pool) AddMember(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.members = append(p.members, user) 
	return nil
}

func (p *pool) RemoveMember(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	members, ok := utils.FindAndRemove(p.members, user)
	if !ok {
		return fmt.Errorf("User is not a member of this pool -- userID: %v; poolID: %v", user, p.id)
	}
	p.members = members
	return nil
}



