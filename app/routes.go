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
	infrastructure.NewPostgreSQL()
	userRep := infrastructure.NewUserRepository()
	levelRep := infrastructure.NewLevelRepository()

	handlers.NewCreateUser(userRep)
	handlers.NewGetUser(userRep)

	handlers.NewGetLevel(levelRep)

	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), handlers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), handlers.GetUser).Methods(http.MethodGet)

	r.HandleFunc(fmt.Sprintf("%s/level", rootPath), handlers.GetLevel).Methods(http.MethodGet)

	r.HandleFunc(fmt.Sprintf("%s/healthCheck", rootPath), handlers.HealthCheck).Methods(http.MethodGet)
}
