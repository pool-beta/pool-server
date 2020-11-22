package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pool-beta/pool-server/daemon/handlers/models"
)

/* User Specific Handlers */

// Creates a new tank for the User
func (h *handler) CreateTank(w http.ResponseWriter, req *http.Request) {
	pools, err := h.pools()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := h.users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request models.RequestCreatePool
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := users.GetUser(request.UserAuth.UserName, request.UserAuth.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tank, err := pools.CreateTankPool(user, request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := &models.ResponseCreatePool{
		UserName: user.UserName(),
		Type:     tank.Type(),
		Name:     tank.Name(),
		ID:       tank.ID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Creates a new pool for the user
func (h *handler) CreatePool(w http.ResponseWriter, req *http.Request) {
	pools, err := h.pools()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := h.users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request models.RequestCreatePool
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := users.GetUser(request.UserAuth.UserName, request.UserAuth.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pool, err := pools.CreatePool(user, request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := &models.ResponseCreatePool{
		UserName: user.UserName(),
		Type:     pool.Type(),
		Name:     pool.Name(),
		ID:       pool.ID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Creates a new drain for the user
func (h *handler) CreateDrain(w http.ResponseWriter, req *http.Request) {
	pools, err := h.pools()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := h.users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request models.RequestCreatePool
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := users.GetUser(request.UserAuth.UserName, request.UserAuth.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	drain, err := pools.CreateDrainPool(user, request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := &models.ResponseCreatePool{
		UserName: user.UserName(),
		Type:     drain.Type(),
		Name:     drain.Name(),
		ID:       drain.ID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Returns information about the given PoolName
func (h *handler) GetPool(w http.ResponseWriter, req *http.Request) {
	pools, err := h.pools()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := h.users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request models.RequestGetPool
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = users.GetUser(request.UserAuth.UserName, request.UserAuth.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Verify user has access to pool

	pool, err := pools.GetPool(request.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := &models.ResponseGetPool{
		Type: pool.Type(),
		Name: pool.Name(),
		ID:   pool.ID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Returns all tanks owned by user
func (h *handler) GetTanks(w http.ResponseWriter, req *http.Request) {
}

// Returns all pools owned by user
func (h *handler) GetPools(http.ResponseWriter, *http.Request) {

}

// Returns all drains owned by user
func (h *handler) GetDrains(http.ResponseWriter, *http.Request) {

}
