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
		Type:     "tank",
		Name:     tank.Name(),
		ID:       tank.ID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Creates a new pool for the user
func (h *handler) CreatePool(http.ResponseWriter, *http.Request) {

}

// Creates a new drain for the user
func (h *handler) CreateDrain(http.ResponseWriter, *http.Request) {

}

// Returns all tanks owned by user
func (h *handler) GetTanks(http.ResponseWriter, *http.Request) {

}

// Returns all pools owned by user
func (h *handler) GetPools(http.ResponseWriter, *http.Request) {

}

// Returns all drains owned by user
func (h *handler) GetDrains(http.ResponseWriter, *http.Request) {

}
