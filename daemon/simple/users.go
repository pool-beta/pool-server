package simple

import (
	"github.com/pool-beta/pool-server/user"
)

/*
	Implements Simple Users
*/

type users struct {
	uf user.UserFactory
}

func NewUsers() (Users, error) {
	// Start UserFactory
	uf, err := user.NewUserFactory()
	if err != nil {
		return nil, err
	}

	return &users{
		uf: uf,
	}, nil
}