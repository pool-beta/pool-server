package handlers

import (
	"io"
	"net/http"
)

type Handler interface {
	HandlerContext

	// Test Handlers
	TestHandler(http.ResponseWriter, *http.Request)
}

type handler struct {
	*handlerContext
}

func NewHandler() (Handler, error) {
	hc, err := newHandlerContext()
	if err != nil {
		return nil, err
	}

	return &handler{
		handlerContext: hc,
	}, nil
}

func (h *handler) TestHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to POOL\n")
}