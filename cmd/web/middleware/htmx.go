package middleware

import (
	"context"
	"net/http"
)

func HTMX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hxHeader := r.Header.Get("HX-Request")
		c := r.Context()
		if hxHeader != "" {
			c = context.WithValue(c, "isHx", true)
		} else {
			c = context.WithValue(c, "isHx", false)
		}
		next.ServeHTTP(w, r)
	})
}
