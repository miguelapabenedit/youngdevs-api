package middlewares

import (
	"net/http"
)

func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method != http.MethodOptions {
			next.ServeHTTP(rw, r)
		} else {
			rw.WriteHeader(http.StatusOK)
			return
		}
	})
}
