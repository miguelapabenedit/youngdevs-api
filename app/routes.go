package app

import (
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/controller"
	"github/miguelapabenedit/youngdevs-api/app/repository"
	"github/miguelapabenedit/youngdevs-api/app/service"
	"net/http"

	"github.com/gorilla/mux"
)

var rep repository.Repository = repository.NewRepository()
var serv service.Service = service.NewServices(rep)
var controllers controller.Controllers = controller.NewController(serv)

func SetUpPublicRoutes(rootPath string, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), controllers.CreateUserHandler).Methods(http.MethodPost, http.MethodHead)
	r.HandleFunc(fmt.Sprintf("%s/user", rootPath), controllers.GetUserHandler).Methods(http.MethodGet, http.MethodHead)

	r.HandleFunc(fmt.Sprintf("%s/healthCheck", rootPath), controllers.HealthCheckHandler).Methods(http.MethodGet)
}
