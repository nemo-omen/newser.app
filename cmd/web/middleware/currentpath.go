package middleware

import (
	"context"
	"net/http"
)

func CurrentPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()
		c = context.WithValue(c, "currentPath", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
