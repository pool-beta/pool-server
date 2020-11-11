package simple

import (
	puser "github.com/pool-beta/pool-server/user"
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
	uf, err := puser.NewUserFactory()
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