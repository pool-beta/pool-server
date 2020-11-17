package types

import (
	"math/rand"
)

// UserID
type UserID = uint64

// NewUserID creates new userID
func NewUserID() UserID {
	r := rand.Uint64()
	return r
}

type UserName = string
