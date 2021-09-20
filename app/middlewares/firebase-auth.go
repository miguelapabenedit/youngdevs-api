package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
)

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

		idtoken := r.Header.Get("Authorization")
		fmt.Println(idtoken)

		token, err := client.VerifyIDToken(ctx, idtoken)
		if err != nil {
			fmt.Printf("error verifying ID token: %v\n", err)
		}

		fmt.Println(token)
		next.ServeHTTP(rw, r)
	})
}
