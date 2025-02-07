package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Chain provides a clean way to apply one or more middleware functions to a
// http.Handler.
func Chain(h http.Handler, m ...Middleware) http.Handler {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}
