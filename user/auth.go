package user

import (
	"fmt"

	. "github.com/pool-beta/pool-server/user/types"
)

const NILUSER = UserID(0)

type UserAuth interface {
	CreateUser(UserName, string) (UserID, error)
	GetUser(UserName, string) (UserID, error)
	DeleteUser(UserName, string) error
}

type userAuth struct {
	passwords map[UserName]string
	ids map[UserName]UserID
}

func initUserAuth() (UserAuth, error) {
	// TODO: connect to db

	passwords := make(map[UserName]string)
	ids := make(map[UserName]UserID)

	return &userAuth{
		passwords: passwords,
		ids: ids,
	}, nil
}


func (ua *userAuth) CreateUser(userName UserName, password string) (UserID, error) {
	if _, exists := ua.ids[userName]; exists {
		return NILUSER, fmt.Errorf("Username is already taken -- username: %v", userName)
	}

	if password == "" {
		return NILUSER, fmt.Errorf("Invalid Password -- password: %v", password)
	}
	
	uid := NewUserID()
	ua.passwords[userName] = password
	ua.ids[userName] = uid
	return uid, nil
}

func (ua *userAuth) GetUser(userName UserName, password string) (UserID, error) {
	err := ua.validateUserPassword(userName, password)
	if err != nil {
		return NILUSER, err
	}
	
	uid, ok := ua.ids[userName]
	if !ok {
		return NILUSER, fmt.Errorf("System broken, no uid -- username: %v", userName)
	}

	return uid, nil
}

func (ua *userAuth) DeleteUser(userName UserName, password string) error {
	err := ua.validateUserPassword(userName, password)
	if err != nil {
		return err
	}

	delete(ua.passwords, userName)
	delete(ua.ids, userName)
	return nil
}

func (ua *userAuth) validateUserPassword(userName UserName, password string) error {
	acutal_password, exists := ua.passwords[userName]
	if !exists {
		return fmt.Errorf("User does not exist -- username: %v", userName)
	}

	if password != acutal_password {
		return fmt.Errorf("Incorrect password")
	}

	return nil
}