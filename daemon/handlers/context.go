package handlers

import ()

/* 
	HandlerContext contains the api for handling requests with the right context
*/

type HandlerContext interface {

}

type handlerContext struct {

}

func newHandlerContext() (*handlerContext, error) {
	return &handlerContext{

	}, nil
}