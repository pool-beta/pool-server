package simple

import (	
	. "github.com/pool-beta/pool-server/user/types"
	. "github.com/pool-beta/pool-server/pool/types"
)

/* 
	Simple contains the simple api for working with pools and streams

	All GET/POST requests will come thru here at one point or another if it needs to interact with pools/streams
*/

type Simple interface {
	Pools() (Pools, error) // POOL (Pools)
	Users() (Users, error) // Users

	CleanUp() error
}

// Simple Users
type Users interface {
	CreateUser() (User, error)
	GetUser(UserID) (User, error)
	RemoveUser(UserID) error

	CleanUp() error
}

type User interface {
	ID() UserID

	Tanks()
	Pools()
	Drains()
	Flows()

	CleanUp() error
}

// TODO: A Space would correspond to a user's network
// Possibly needed to scale
type Space interface {
	
}

// Pools implements POOL (simple pools)
type Pools interface {
	CreatePool(User, string) (Pool, error)
	CreateDrainPool(User, string) (Drain, error) 
	CreateTankPool(User, string) (Tank, error)
	GetPool(PoolID) (Pool, error) 
	RemovePool(PoolID) error

	CleanUp() error
}

type Pool interface {
	CreateStream(Pool) (Stream, error)
	GetStream(StreamID) (Stream, error)

	CreatePushFlow() (Flow, error)
	CreatePullFlow() (Flow, error)

	CleanUp() error
}

type Tank interface {
	Pool
}

type Drain interface {
	Pool
}

type Stream interface {
	EnableOverDraft()
	DisableOverDraft()
	EnableFlexibleOverdraft()
	DisableFlexibleOverdraft()
	SetPercentOverdraft()
	SetMaxOverdraft()
	// SetMinOverdraft()
}

// Flow is the api to interact with drops
type Flow interface {
	Accept()
	Decline()
}

// -----------------------------------------------------------------------------------------------------

type simple struct {
	users Users
	pools Pools
}

// Should only be called once
// Start all the necessary Factories
func NewSimple() (Simple, error) {
	// TODO: Start/Connect Database
	users, err := InitUsers()
	if err != nil {
		return nil, err
	}

	pools, err := InitPools()
	if err != nil {
		return nil, err
	}

	return &simple{
		users: users,
		pools: pools,
	}, nil
}

func (s *simple) Pools() (Pools, error) {
	return s.pools, nil
}

func (s *simple) Users() (Users, error) {
	return s.users, nil
}

func (s *simple) CleanUp() error {
	var err error

	err = s.pools.CleanUp()
	if err != nil {
		return err
	}

	err = s.users.CleanUp()
	if err != nil {
		return err
	}
	return nil
}