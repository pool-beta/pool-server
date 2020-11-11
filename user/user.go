package user

import (
	"fmt"
	"sync"

	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/user/types"
)

type UserFactory interface {
	CreateUser(string, USDollar) (User, error)
	RetrieveUser(UserID) (User, error)
	ReturnUser(UserID) error
}

type userRef struct {
	user User
	refCount int
	mutex sync.Mutex
}

type userFactory struct {
	users map[UserID]*userRef
}

func NewUserFactory() (UserFactory, error) {
	users := make(map[UserID]*userRef)
	return &userFactory{
		users: users,
	}, nil
}

func (uf *userFactory) CreateUser(userName string, amount USDollar) (User, error) {
	// Create a new user id
	uid := NewUserID()

	user := initUser(uid, userName, amount)

	// Create poolRef
	ur := &userRef{
		user: user,
		refCount: 1,
	}

	uf.users[uid] = ur
	return user, nil
}

func (uf *userFactory) RetrieveUser(uid UserID) (User, error) {
	ur, ok := uf.users[uid]
	if !ok {
		return nil, fmt.Errorf("User does not exist -- user: %v", uid)
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	ur.refCount++
	return ur.user, nil
}

func (uf *userFactory) ReturnUser(uid UserID) error {
	ur, ok := uf.users[uid]
	if !ok {
		return fmt.Errorf("User does not exist -- user: %v", uid)
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	ur.refCount--

	// TODO: Flush when refCount == 0
	return nil
}

// -----------------------------------------------------------------------------------------------------------

type User interface {
	// Returns the current amount in the reserve
	GetReserve() USDollar
	// Puts money into the user's reserve
	Deposit(USDollar) error
}

type user struct {
	id UserID
	name string
	reserve USDollar

	mutex sync.Mutex // Need for multiple logins
}

func initUser(id UserID, name string, amount USDollar) *user {
	return &user{
		id: id,
		name: name,
		reserve: amount,
	}
}

func (u *user) GetReserve() USDollar {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.reserve
}

func (u *user) Deposit(amount USDollar) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	if amount < 0 {
		return fmt.Errorf("Can't deposit a negative amount of money -- amount: %v", amount)
	}

	u.reserve += amount
	return nil
}