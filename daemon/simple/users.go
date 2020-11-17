package simple

import (
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/types"
	puser "github.com/pool-beta/pool-server/user"
	. "github.com/pool-beta/pool-server/user/types"
)

/*
	Implements Simple Users
*/

type users struct {
	pools Pools
	uf    puser.UserFactory
}

type user struct {
	pools Pools
	user  puser.User
}

func InitUsers(pools Pools) (Users, error) {
	// Start UserFactory
	uf, err := puser.InitUserFactory()
	if err != nil {
		return nil, err
	}

	return &users{
		pools: pools,
		uf:    uf,
	}, nil
}

func (us *users) CleanUp() error {
	return nil
}

func (us *users) CreateUser(name UserName, password string, amount USDollar) (User, error) {
	u, err := us.uf.CreateUser(name, password, amount)
	if err != nil {
		return nil, err
	}

	return &user{
		pools: us.pools,
		user:  u,
	}, nil
}

func (us *users) GetUser(name UserName, password string) (User, error) {
	u, err := us.uf.RetrieveUser(name, password)
	if err != nil {
		return nil, err
	}

	return &user{
		pools: us.pools,
		user:  u,
	}, nil
}

func (us *users) RemoveUser(name UserName, password string) error {
	err := us.uf.RemoveUser(name, password)
	return err
}

// Testing
func (us *users) GetAllUserNames() ([]UserName, error) {
	return us.uf.RetreieveAllUserNames()
}

func (u *user) ID() UserID {
	return u.user.GetID()
}

func (u *user) UserName() UserName {
	return u.user.GetUserName()
}

func (u *user) AddTank(pid PoolID) error {
	return u.user.AddTank(pid)
}

func (u *user) Tanks() ([]Tank, error) {
	pids := u.user.GetTanks()

	// Create simple tanks to return to return
	tanks := make([]Tank, len(pids))

	for i, pid := range pids {
		tank, err := u.pools.GetPool(pid)
		if err != nil {
			return nil, err
		}

		tanks[i] = tank
	}

	return tanks, nil
}

func (u *user) AddPool(pid PoolID) error {
	return u.user.AddPool(pid)
}

func (u *user) Pools() ([]Pool, error) {
	pids := u.user.GetPools()

	// Create simple tanks to return to return
	pools := make([]Pool, len(pids))

	for i, pid := range pids {
		pool, err := u.pools.GetPool(pid)
		if err != nil {
			return nil, err
		}

		pools[i] = pool
	}

	return pools, nil
}

func (u *user) AddDrain(pid PoolID) error {
	return u.user.AddDrain(pid)
}

func (u *user) Drains() ([]Drain, error) {
	pids := u.user.GetTanks()

	// Create simple tanks to return to return
	drains := make([]Drain, len(pids))

	for i, pid := range pids {
		drain, err := u.pools.GetPool(pid)
		if err != nil {
			return nil, err
		}

		drains[i] = drain
	}

	return drains, nil
}

func (u *user) Flows() []Flow {
	return nil
}

func (u *user) CleanUp() error {
	return nil
}
