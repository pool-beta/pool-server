package user

import (
	. "github.com/pool-beta/pool-server/types"
)

type User interface {

}

type user struct {
	id UserID
	name string
}