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
	levelStateRepo := infrastructure.NewLevelStateRepository()

	handlers.NewCreateUser(userRep)
	handlers.NewGetAllUsers(userRep)
	handlers.NewGetUser(userRep)
	handlers.NewUpdateUser(userRep)
	handlers.NewGetRanking(userRep)
	handlers.NewGetLevelState(levelStateRepo)
	handlers.NewUpdateLevelState(levelStateRepo)
	handlers.NewGetAllLevelState(levelStateRepo)
	handlers.NewValidateLevel(levelStateRepo, levelRep)

	handlers.NewGetLevel(levelRep)

	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), handlers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), handlers.GetUser).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), handlers.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc(fmt.Sprintf("%s/user/reset/{id}", rootPath), handlers.ResetUser).Methods(http.MethodDelete)
	r.HandleFunc(fmt.Sprintf("%s/users", rootPath), handlers.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/users/ranking", rootPath), handlers.GetRanking).Methods(http.MethodGet)

	r.HandleFunc(fmt.Sprintf("%s/level", rootPath), handlers.GetLevel).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/levels", rootPath), handlers.GetAllLevels).Methods(http.MethodGet)

	r.HandleFunc(fmt.Sprintf("%s/level/state/{level}", rootPath), handlers.GetLevelState).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/level/states", rootPath), handlers.GetAllLevelState).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/level/state", rootPath), handlers.UpdateLevelState).Methods(http.MethodPut)
	r.HandleFunc(fmt.Sprintf("%s/level/validate", rootPath), handlers.ValidateLevel).Methods(http.MethodPost)

	r.HandleFunc(fmt.Sprintf("%s/healthCheck", rootPath), handlers.HealthCheck).Methods(http.MethodGet)
}
