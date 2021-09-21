package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
)

var env = os.Getenv("ENV")

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		fbapp, err := firebase.NewApp(ctx, nil)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
		client, err := fbapp.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		if env != "DEV" {

			if len(splitToken) != 2 || splitToken[1] == "" {
				fmt.Println("No autherization token provided")
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}

			token, err := client.VerifyIDToken(ctx, splitToken[1])
			if err != nil {
				fmt.Printf("error verifying ID token: %v\n", err)
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}

			fmt.Println(token)
		} else {
			fmt.Println("Recieved token:", splitToken)
		}

		next.ServeHTTP(rw, r)
	})
}
