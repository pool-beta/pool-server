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
	GetTanks() []PoolID
	GetPools() []PoolID
	GetDrains() []PoolID

	AddTank(PoolID) error
	AddPool(PoolID) error
	AddDrain(PoolID) error
}

type user struct {
	id      UserID
	name    string
	reserve USDollar

	tanks  []PoolID
	pools  []PoolID
	drains []PoolID

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

func (u *user) GetTanks() []PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.tanks
}

func (u *user) AddTank(pid PoolID) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	// TODO: validate pid

	u.tanks = append(u.tanks, pid)
	return nil
}

func (u *user) GetPools() []PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.pools
}

func (u *user) AddPool(pid PoolID) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	// TODO: validate pid

	u.pools = append(u.pools, pid)
	return nil
}

func (u *user) AddDrain(pid PoolID) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	// TODO: validate pid

	u.drains = append(u.drains, pid)
	return nil
}

func (u *user) GetDrains() []PoolID {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	return u.drains
}
