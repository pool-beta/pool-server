package handlers

import (
	"github.com/pool-beta/pool-server/daemon/simple"
)

/* 
	HandlerContext contains the api for handling requests with the right context
*/

type HandlerContext interface {

}

type handlerContext struct {

}

func newHandlerContext() (*handlerContext, error) {
	s := simple.NewSimple()

	return &handlerContext{
		simple: s,
	}, nil
}