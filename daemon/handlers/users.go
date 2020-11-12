package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pool-beta/pool-server/daemon/handlers/models"
)

func (h *handler) CreateUser(w http.ResponseWriter, req *http.Request) {
	simple, err := h.Simple()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usrs, err := simple.Users()
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
		UserID: user.ID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUsers(w http.ResponseWriter, req *http.Request) {
	simple, err := h.Simple()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usrs, err := simple.Users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userNames, err := usrs.GetAllUserNames()

	resp := make([]models.User, 0)
	for i := 0; i < len(userNames); i++ {
		resp = append(resp, models.User{
			UserName: userNames[i],
		})
	}
	
	users := models.Users{
		Users: resp,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}