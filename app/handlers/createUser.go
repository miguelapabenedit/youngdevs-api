package handlers

import (
	"encoding/json"
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/data"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"io/ioutil"
	"net/http"
)

var rep repository.CreateUser

func NewCreateUser(createUserRepository repository.CreateUser) {
	rep = createUserRepository
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	bodyContent, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	err = json.Unmarshal(bodyContent, user)

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rep.CreateUser(user)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contet-Type", "appcontent/json")
	w.Write(uj)
}
