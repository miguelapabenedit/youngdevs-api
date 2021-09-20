package handlers

import (
	"encoding/json"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"net/http"
)

var getUserRepository repository.GetUser

func NewGetUser(userRepository repository.GetUser) {
	getUserRepository = userRepository
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "Debe enviar id", http.StatusBadRequest)
		return
	}

	user := getUserRepository.GetUser(id)
	msg, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
