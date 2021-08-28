package controller

import (
	"encoding/json"
	"net/http"
)

type controllers struct{}

// para simpleza vamos a manejar todo con una sola interfaz
type Controllers interface {
	HealthCheckHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
}

func NewController() Controllers {
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
	w.WriteHeader(http.StatusNotImplemented)
}
func (c *controllers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
