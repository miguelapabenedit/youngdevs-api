package handlers

import (
	"encoding/json"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"net/http"
)

var getAllLevelStateRepository repository.GetAllUserLevelState

func NewGetAllLevelState(repo repository.GetAllUserLevelState) {
	getAllLevelStateRepository = repo
}

func GetAllLevelState(w http.ResponseWriter, r *http.Request) {
	authUserId := r.Header.Get("AuthProviderUserId")

	user := userRepository.Get(authUserId)

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uls := getAllLevelStateRepository.GetAllUserLevelState(int(user.ID))

	msg, err := json.Marshal(&uls)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
