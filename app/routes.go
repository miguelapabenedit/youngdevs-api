package app

import (
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/handlers"
	"github/miguelapabenedit/youngdevs-api/app/infrastructure"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpPublicRoutes(rootPath string, r *mux.Router, l *log.Logger) {
	rep := infrastructure.NewPostgreSQL()
	handlers.NewCreateUser(rep)
	handlers.NewGetUser(rep)

	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), handlers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), handlers.GetUser).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/healthCheck", rootPath), handlers.HealthCheck).Methods(http.MethodGet)
}
