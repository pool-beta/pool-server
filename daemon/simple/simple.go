package simple

import (

)

/* 
	Simple contains the simple api for working with pools and streams

	All GET/POST requests will come thru here at one point or another if it needs to interact with pools/streams
*/

type Simple interface {
	// TODO: Cleanup()

	Pools() (Pools, error) // POOL
	Users() (Users, error) // Users
}

// Simple Users
type Users interface {
	NewUser()
	GetUser()
	RemoveUser()
}

type User interface {
	
}

type Space interface {

}

// Pools implements POOL (simple pools)
type Pools interface {
	NewPool() Pool
	NewDrainPool() Drain 
	NewTankPool() Tank
	RemovePool() error

	NewStream() Stream
	RemoveStream() error
}

type Pool interface {
	NewFlow() Flow
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
	users, err := NewUsers()
	if err != nil {
		return nil, err
	}

	pools, err := NewPools()
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