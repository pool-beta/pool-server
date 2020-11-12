package handlers

import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
	
	. "github.com/pool-beta/pool-server/daemon/handlers/models"
)

type Handler interface {
	HandlerContext

	// Test Handlers
	TestHandler(http.ResponseWriter, *http.Request)

	// Users
	CreateUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
}

type handler struct {
	*handlerContext
}

/* Should only be called once */
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
	var test Test

	err := json.NewDecoder(req.Body).Decode(&test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	io.WriteString(w, fmt.Sprintf("Welcome to POOL\ntest: %v\n", test.Test))
}