package controller

import (
	"encoding/json"
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/entity"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"github/miguelapabenedit/youngdevs-api/app/service"
	"io/ioutil"
	"net/http"
)

type controllers struct{}

var serv service.Service

// para simpleza vamos a manejar todo con una sola interfaz
type Controllers interface {
	HealthCheckHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
}

func NewController(service service.Service) Controllers {
	serv = service
	return &controllers{}
}

func (c *controllers) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := json.Marshal("Im alive Belen")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

func (c *controllers) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "Debe enviar id", http.StatusBadRequest)
		return
	}

	user := repository.GetUser(id)
	msg, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(msg)
}

func (c *controllers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &entity.User{}
	bodyContent, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bodyContent, user)

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = serv.CreateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(userJson)
	w.Header().Set("Contet-Type", "appcontent/json")
}
