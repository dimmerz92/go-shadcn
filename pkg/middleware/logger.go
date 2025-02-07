package middleware

import (
	"log"
	"net/http"
	"os"
)

// Logger is a simple middleware that logs the method and URL of a request to
// the standard output.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetOutput(os.Stdout)
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
