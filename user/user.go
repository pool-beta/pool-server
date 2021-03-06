package user

import (
	"fmt"
	"sync"

	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/user/types"
)

// Factory is used to create/remove users
type UserFactory interface {
	CreateUser(UserName, string, USDollar) (User, error)
	RetrieveUser(UserName, string) (User, error)
	ReturnUser(UserID) error
	RemoveUser(UserName, string) error

	// Testing
	RetreieveAllUserNames() ([]UserName, error)
}

type userRef struct {
	user     User
	refCount int
	mutex    sync.Mutex
}

type userFactory struct {
	userAuth UserAuth
	users    map[UserID]*userRef
}

// InitFactory initializes the factory for users
func InitUserFactory() (UserFactory, error) {
	// Init UserAuth
	userAuth, err := initUserAuth()
	if err != nil {
		return nil, err
	}

	users := make(map[UserID]*userRef)
	return &userFactory{
		userAuth: userAuth,
		users:    users,
	}, nil
}

func (uf *userFactory) CreateUser(userName UserName, password string, amount USDollar) (User, error) {
	// Create a new user id thru userAuth
	uid, err := uf.userAuth.CreateUser(userName, password)
	if err != nil {
		return nil, err
	}

	user := initUser(uid, userName, amount)

	// Create poolRef
	ur := &userRef{
		user:     user,
		refCount: 1,
	}

	uf.users[uid] = ur
	return user, nil
}

func (uf *userFactory) RetrieveUser(userName UserName, password string) (User, error) {
	// Get uid thru userAuth
	uid, err := uf.userAuth.GetUser(userName, password)
	if err != nil {
		return nil, err
	}

	ur, ok := uf.users[uid]
	if !ok {
		return nil, fmt.Errorf("User does not exist -- username: %v", userName)
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

func (uf *userFactory) RemoveUser(userName UserName, password string) error {
	return uf.userAuth.DeleteUser(userName, password)
}

func (uf *userFactory) RetreieveAllUserNames() ([]UserName, error) {
	return uf.userAuth.GetAllUserNames()
}

// -----------------------------------------------------------------------------------------------------------

// User is the interface for working with a User
type User interface {
	GetID() UserID
	GetUserName() UserName
	// Returns the current amount in the reserve
	GetReserve() USDollar
	// Puts money into the user's reserve
	Deposit(USDollar) error

	// Pools
	GetTank(string) PoolID
	GetPool(string) PoolID
	GetDrain(string) PoolID

	AddTank(string, PoolID) error
	AddPool(string, PoolID) error
	AddDrain(string, PoolID) error

	GetTanks() map[string]PoolID
	GetPools() map[string]PoolID
	GetDrains() map[string]PoolID
}

type user struct {
	id      UserID
	name    string
	reserve USDollar

	tanks  map[string]PoolID
	pools  map[string]PoolID
	drains map[string]PoolID

	mutex sync.Mutex // Need for multiple logins
}

func initUser(id UserID, name string, amount USDollar) *user {
	return &user{
		id:      id,
		name:    name,
		reserve: amount,
	}
}

func (u *user) GetID() UserID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.id
}

func (u *user) GetUserName() UserName {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.name
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

func (u *user) GetTank(name string) PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.tanks[name]
}

func (u *user) GetTanks() map[string]PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.tanks
}

func (u *user) AddTank(name string, pid PoolID) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	// TODO: validate pid
	if _, exists := u.tanks[name]; exists {
		return fmt.Errorf("Name not avaiable -- name: %v", name)
	}
	u.tanks[name] = pid
	return nil
}

func (u *user) GetPool(name string) PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.pools[name]
}

func (u *user) GetPools() map[string]PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.pools
}

func (u *user) AddPool(name string, pid PoolID) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	// TODO: validate pid
	if _, exists := u.pools[name]; exists {
		return fmt.Errorf("Name not avaiable -- name: %v", name)
	}
	u.pools[name] = pid
	return nil
}

func (u *user) GetDrain(name string) PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.drains[name]
}

func (u *user) GetDrains() map[string]PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.drains
}

func (u *user) AddDrain(name string, pid PoolID) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	// TODO: validate pid
	if _, exists := u.drains[name]; exists {
		return fmt.Errorf("Name not avaiable -- name: %v", name)
	}
	u.drains[name] = pid
	return nil
}
