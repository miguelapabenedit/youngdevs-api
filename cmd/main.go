package main

import (
	"fmt"
	"github/miguelapabenedit/youngdevs-api/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const apiRootPath string = "/api"

func main() {

	fmt.Println("Stargin Server at port:8000")
	r := mux.NewRouter()

	app.SetUpPublicRoutes(apiRootPath, r)
	// start the server in a go function with an escape logic
	log.Fatalln(http.ListenAndServe(":3000", r))
}
