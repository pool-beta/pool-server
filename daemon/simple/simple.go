package simple

import (

)

/* 
	Simple contains the simple api for working with pools and streams

	All GET/POST requests will come thru here at one point or another if it needs to interact with pools/streams
*/

type Simple interface {
	Pools() (Pools, error) // POOL
	Users() (Users, error) // Users
}

type Users interface {
	CreateUser()
}

// Pools implements POOL
type Pools interface {
	CreatePool()
}
