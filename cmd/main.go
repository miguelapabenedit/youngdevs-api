package main

import (
	"context"
	"github/miguelapabenedit/youngdevs-api/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const apiRootPath string = "/api"

var (
	port string = os.Getenv("PORT")
)

func main() {

	l := log.New(os.Stdout, "youngdevs-api: ", log.LstdFlags)
	r := mux.NewRouter()

	app.SetUpPublicRoutes(apiRootPath, r)

	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"POST, GET, OPTIONS, PUT, DELETE"}),
		gohandlers.AllowedHeaders([]string{"Content-Type"}),
	)

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      ch(r),
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 120 * time.Second,
	}
	// start the server in a go function with an escape logic
	go func() {
		l.Println("Starting server on port" + port)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error stargin server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := <-c

	log.Println("Got signal:", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
