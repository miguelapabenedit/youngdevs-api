package app

import (
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app/controller"
	"net/http"

	"github.com/gorilla/mux"
)

var controllers controller.Controllers = controller.NewController()

func SetUpPublicRoutes(rootPath string, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("%s/healthCheck", rootPath), controllers.HealthCheckHandler).Methods(http.MethodGet)
}
