package handlers

import (
	"fmt"

	"github.com/pool-beta/pool-server/daemon/simple"
)

/*
	HandlerContext contains the api for handling requests with the right context
*/

type HandlerContext interface {
	simple() (simple.Simple, error)
	pools() (simple.Pools, error)
	users() (simple.Users, error)
}

type handlerContext struct {
	simp simple.Simple
}

/* Should only be called once */
func newHandlerContext() (*handlerContext, error) {
	s, err := simple.NewSimple()
	if err != nil {
		return nil, err
	}

	return &handlerContext{
		simp: s,
	}, nil
}

func (hc *handlerContext) pools() (simple.Pools, error) {
	s, err := hc.simple()
	if err != nil {
		return nil, err
	}

	p, err := s.Pools()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (hc *handlerContext) users() (simple.Users, error) {
	s, err := hc.simple()
	if err != nil {
		return nil, err
	}

	u, err := s.Users()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (hc *handlerContext) simple() (simple.Simple, error) {
	if hc.simp == nil {
		return nil, fmt.Errorf("simple is nil in handlerContext")
	}
	return hc.simp, nil
}
