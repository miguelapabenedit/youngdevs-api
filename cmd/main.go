package main

import (
	"github/miguelapabenedit/youngdevs-api/app"
	"log"
	"os"

	"github.com/gorilla/mux"
)

const apiRootPath string = "/api"

var (
	port string = os.Getenv("PORT")
)

func main() {
	l := log.New(os.Stdout, "youngdevs-api: ", log.LstdFlags)
	r := mux.NewRouter()

	app.SetUpPublicRoutes(apiRootPath, r, l)

	app.StartServer(r, l, port)
}
