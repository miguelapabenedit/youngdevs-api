package app

import (
	"context"
	"github/miguelapabenedit/youngdevs-api/app/middlewares"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func StartServer(r *mux.Router, l *log.Logger, port string) {
	if port == "" {
		port = "3030"
	}

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      middlewares.CORSHandler(middlewares.AuthHandler(r)),
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
