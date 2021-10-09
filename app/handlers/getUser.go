package handlers

import (
	"encoding/json"
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"net/http"
)

var userRepository repository.User

func NewGetUser(repo repository.User) {
	userRepository = repo
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	authUserId := r.Header.Get("AuthProviderUserId")
	email := r.Header.Get("Email")

	if len(authUserId) < 1 || len(email) < 1 {
		fmt.Println("No se encontro id o email en la información del usuario")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := userRepository.Get(authUserId)

	if user.ID == 0 {
		user.AuthProviderUserId = authUserId
		user.Email = email
		err := userRepository.Create(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user = userRepository.Get(authUserId)
	}

	msg, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
