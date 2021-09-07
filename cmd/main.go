package main

import (
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const apiRootPath string = "/api"

var (
	port string = os.Getenv("PORT")
)

func main() {

	fmt.Println("Stargin Server at port:" + port)
	r := mux.NewRouter()

	app.SetUpPublicRoutes(apiRootPath, r)

	ch := handlers.CORS(
		handlers.AllowedOrigins([]string{"https://youngdevs-e5ff0.web.app/"}),
	)(r)

	// start the server in a go function with an escape logic
	log.Fatalln(http.ListenAndServe(":"+port, ch))
}
