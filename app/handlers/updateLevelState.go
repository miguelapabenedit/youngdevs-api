package handlers

import (
	"encoding/json"
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"io/ioutil"
	"net/http"
)

var updateLevelStateRepository repository.UpdateUserLevelState

func NewUpdateLevelState(updateRepository repository.UpdateUserLevelState) {
	updateLevelStateRepository = updateRepository
}

func UpdateLevelState(w http.ResponseWriter, r *http.Request) {
	uls := &data.UserLevelState{}
	bodyContent, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bodyContent, uls)

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authUserId := r.Header.Get("AuthProviderUserId")
	user := userRepository.Get(authUserId)

	uls.UserID = user.ID

	if uls.UserID == 0 || uls.LevelID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = updateLevelStateRepository.UpdateLevelState(uls)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(uls)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contet-Type", "appcontent/json")
	w.Write(uj)
}
