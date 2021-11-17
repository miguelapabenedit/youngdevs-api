package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ResetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := userRepository.GetById(id)

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Score = 0
	user.CurrentLevel = 1

	err := levelStateRepo.DeleteAllById(user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userRepository.Update(user)

	w.WriteHeader(http.StatusOK)
}
