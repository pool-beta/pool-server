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

func (pFactory *poolFactory) CreatePool(owner UserID, poolName string) (Pool, error) {
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
	AdminCheck(user UserID, level string) bool
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

func (p *pool) AddPusher(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.pushers = append(p.pushers, user) 
	return nil
}

func (p *pool) RemovePusher(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	pushers, ok := utils.FindAndRemove(p.pushers, user)
	if !ok {
		return fmt.Errorf("User is not a pusher of this pool -- userID: %v; poolID: %v", user, p.id)
	}
	p.pushers = pushers
	return nil
}

func (p *pool) AddPuller(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.pullers = append(p.pullers, user) 
	return nil
}

func (p *pool) RemovePuller(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	pullers, ok := utils.FindAndRemove(p.pullers, user)
	if !ok {
		return fmt.Errorf("User is not a puller of this pool -- userID: %v; poolID: %v", user, p.id)
	}
	p.pullers = pullers
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
	p.pushers = admins
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

// Triggers

func (p *pool) Push() error {
	return nil
}

func (p *pool) Pull() error {
	return nil
}

func (p *pool) Reset() error {
	return nil
}


