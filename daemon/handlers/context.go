package handlers

import (
	"fmt"

	"github.com/pool-beta/pool-server/daemon/simple"
)

/* 
	HandlerContext contains the api for handling requests with the right context
*/

type HandlerContext interface {
	Simple() (simple.Simple, error)
}

type handlerContext struct {
	simple simple.Simple
}

/* Should only be called once */
func newHandlerContext() (*handlerContext, error) {
	s, err := simple.NewSimple()
	if err != nil {
		return nil, err
	}

	return &handlerContext{
		simple: s,
	}, nil
}

func (hc *handlerContext) Simple() (simple.Simple, error) {
	if hc.simple == nil {
		return nil, fmt.Errorf("simple is nil in handlerContext")
	}
	return hc.simple, nil
}