package handlers

import (
	"encoding/json"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var userLevelStateRepo repository.UserLevelState

func NewGetLevelState(repo repository.UserLevelState) {
	userLevelStateRepo = repo
}

func GetLevelState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authUserId := r.Header.Get("AuthProviderUserId")

	user := userRepository.Get(authUserId)

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sLevel, ok := vars["level"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lvl, err := strconv.Atoi(sLevel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userLevelState := &data.UserLevelState{UserID: user.ID, LevelID: uint(lvl)}

	err = userLevelStateRepo.GetLevelState(userLevelState)

	if err != nil && levelRepo.Exist(userLevelState.LevelID) {
		if err = userLevelStateRepo.CreateLevelState(userLevelState); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	msg, err := json.Marshal(userLevelState)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}
