package pool

import (
	"fmt"
	"sync"

	. "github.com/pool-beta/pool-server/types"
	"github.com/pool-beta/pool-server/utils"
)

type PoolAuth interface {
	// Admin Control
	AdminCheck(user UserID, level string) bool
	AddOwner(UserID) error
	RemoveOwner(UserID) error
	AddAdmin(UserID) error
	RemoveAdmin(UserID) error
	AddMember(UserID) error
	RemoveMember(UserID) error
}

type poolAuth struct {
	pid PoolID // For logging

	owners []UserID
	admins []UserID
	members []UserID

	mutex sync.Mutex // Not Needed if Pool is lock 
}

func NewPoolAuth(pid PoolID) *poolAuth {
	owners := make([]UserID, 1)
	admins := make([]UserID, 0)
	members := make([]UserID, 0)

	return &poolAuth{
		pid: pid,
		owners: owners,
		admins: admins,
		members: members,
	}
}
// Admin
func (p *poolAuth) AdminCheck(user UserID, level string) bool {
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

func (p *poolAuth) AddOwner(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.owners = append(p.owners, user) 
	return nil
}

func (p *poolAuth) RemoveOwner(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	owners, ok := utils.FindAndRemove(p.owners, user)
	if !ok {
		return fmt.Errorf("User is not an owner of this pool -- userID: %v; poolID: %v", user, p.pid)
	}
	p.owners = owners
	return nil
}

func (p *poolAuth) AddAdmin(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.admins = append(p.admins, user) 
	return nil
}

func (p *poolAuth) RemoveAdmin(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	admins, ok := utils.FindAndRemove(p.admins, user)
	if !ok {
		return fmt.Errorf("User is not an admin of this pool -- userID: %v; poolID: %v", user, p.pid)
	}
	p.admins = admins
	return nil
}

func (p *poolAuth) AddMember(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.members = append(p.members, user) 
	return nil
}

func (p *poolAuth) RemoveMember(user UserID) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	members, ok := utils.FindAndRemove(p.members, user)
	if !ok {
		return fmt.Errorf("User is not a member of this pool -- userID: %v; poolID: %v", user, p.pid)
	}
	p.members = members
	return nil
}
