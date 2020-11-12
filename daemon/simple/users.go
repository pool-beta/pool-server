package simple

import (
	puser "github.com/pool-beta/pool-server/user"
	. "github.com/pool-beta/pool-server/user/types"
	. "github.com/pool-beta/pool-server/types"
)

/*
	Implements Simple Users
*/

type users struct {
	uf puser.UserFactory
}

type user struct {
	user puser.User
}

func InitUsers() (Users, error) {
	// Start UserFactory
	uf, err := puser.InitUserFactory()
	if err != nil {
		return nil, err
	}

	return &users{
		uf: uf,
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
		user: u,
	}, nil
}

func (us *users) GetUser(name UserName, password string) (User, error) {
	u, err := us.uf.RetrieveUser(name, password)
	if err != nil {
		return nil, err
	}

	return &user{
		user: u,
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

func (u *user) Tanks() []Tank {
	return nil
}

func (u *user) Pools() []Pool {
	return nil
}

func (u *user) Drains() []Drain {
	return nil
}

func (u *user) Flows() []Flow {
	return nil
}

func (u *user) CleanUp() error {
	return nil
}
