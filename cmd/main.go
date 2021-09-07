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

	handler := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r)
	// start the server in a go function with an escape logic
	log.Fatalln(http.ListenAndServe(":"+port, handler))
}
