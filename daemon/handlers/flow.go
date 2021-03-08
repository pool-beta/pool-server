package handlers

import (
	// "encoding/json"
	"net/http"

	// "github.com/pool-beta/pool-server/daemon/handlers/models"
)

func (h *handler) PayDebit(w http.ResponseWriter, req *http.Request) {
	// var err error

	// // Decode Request
	// var request models.RequestPayDebit
	// err = json.NewDecoder(req.Body).Decode(&request)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// users, err := h.users()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // Get user
	// user, err := users.GetUser(request.UserAuth.UserName, request.UserAuth.Password)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // Get PoolID
	// pools, err := h.pools()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// pool, err := pools.GetPool(request.ID)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

}
