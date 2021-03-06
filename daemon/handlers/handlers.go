package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pool-beta/pool-server/daemon/handlers/models"
	"github.com/pool-beta/pool-server/daemon/simple"
)

type Handler interface {
	HandlerContext

	// Test Handlers
	TestHandler(http.ResponseWriter, *http.Request)
	TestSetup(http.ResponseWriter, *http.Request)
	TestReset(http.ResponseWriter, *http.Request)

	// Users
	CreateUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)

	// User
	CreateTank(http.ResponseWriter, *http.Request)
	CreatePool(http.ResponseWriter, *http.Request)
	CreateDrain(http.ResponseWriter, *http.Request)

	GetPool(http.ResponseWriter, *http.Request)

	GetTanks(http.ResponseWriter, *http.Request)
	GetPools(http.ResponseWriter, *http.Request)
	GetDrains(http.ResponseWriter, *http.Request)

	// Flow (Push/Pull)
	PayDebit(http.ResponseWriter, *http.Request)
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
	var test models.Test

	err := json.NewDecoder(req.Body).Decode(&test)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	io.WriteString(w, fmt.Sprintf("Welcome to POOL\ntest: %v\n", test.Test))
}

func (h *handler) TestSetup(w http.ResponseWriter, req *http.Request) {
	var err error

	SECRET := "admin"
	DEF_PASSWORD := "test"

	// Decode Request
	var request models.RequestSetupTestEnv
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate Secret
	if request.Secret != SECRET {
		http.Error(w, "Wrong Password -- Fuck Off!", http.StatusBadRequest)
		return
	}

	/* Setup Users */
	usrs, err := h.users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alice, err := usrs.CreateUser("alice", DEF_PASSWORD, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bob, err := usrs.CreateUser("bob", DEF_PASSWORD, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Consolidate Users
	users := []simple.User{
		alice,
		bob,
	}

	/* Create Tanks/Pools/Drains */
	pools, err := h.pools()

	for _, u := range users {
		tank, err := pools.CreateTankPool(u, "tank0")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pool, err := pools.CreatePool(u, "pool0")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		drain, err := pools.CreateDrainPool(u, "drain0")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add Streams
		_, err = tank.CreateStream(pool)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = pool.CreateStream(drain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	/* Create Response */
	// User Model
	var modelUsers []models.TestUser
	for _, u := range users {
		modelUsers = append(modelUsers, models.TestUser{
			UserName: u.UserName(),
			Password: DEF_PASSWORD,
			UserID:   u.ID(),
		})
	}

	response := &models.ResponseSetupTestEnv{
		Users: modelUsers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *handler) TestReset(w http.ResponseWriter, req *http.Request) {
	var err error

	SECRET := "admin"

	// Decode Request
	var request models.RequestSetupTestEnv
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate Secret
	if request.Secret != SECRET {
		http.Error(w, "Wrong Password -- Fuck Off!", http.StatusBadRequest)
		return
	}

	// TODO: Reset more gracefully

	hc, err := newHandlerContext()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.handlerContext = hc
	w.WriteHeader(http.StatusOK)
}
