package simple

import (
	puser "github.com/pool-beta/pool-server/user"
	. "github.com/pool-beta/pool-server/user/types"
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

func (us *users) CreateUser() (User, error) {
	return nil, nil
}

func (us *users) GetUser(UserID) (User, error) {
	return nil, nil
}
	
func (us *users) RemoveUser(UserID) error {
	return nil
}
