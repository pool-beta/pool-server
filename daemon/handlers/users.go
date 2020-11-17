package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pool-beta/pool-server/daemon/handlers/models"
)

// Creates a new user
func (h *handler) CreateUser(w http.ResponseWriter, req *http.Request) {
	usrs, err := h.users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request models.RequestCreateUser
	err = json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usrs.CreateUser(request.UserName, request.Password, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := &models.ResponseCreateUser{
		UserName: user.UserName(),
		UserID:   user.ID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Returns all users
func (h *handler) GetUsers(w http.ResponseWriter, req *http.Request) {
	usrs, err := h.users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	allUserNames, err := usrs.GetAllUserNames()

	userNames := make([]models.UserName, 0)
	for i := 0; i < len(allUserNames); i++ {
		userNames = append(userNames, models.UserName{
			UserName: allUserNames[i],
		})
	}

	resp := models.ResponseGetUsers{
		UserNames: userNames,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
