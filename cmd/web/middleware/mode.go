package middleware

import (
	"context"
	"net/http"
)

type Mode struct {
	Dev  bool
	Prod bool
}

func NewMode(mode string) *Mode {
	if mode == "development" {
		return &Mode{Dev: true, Prod: false}
	}
	return &Mode{Dev: false, Prod: true}
}

func (m *Mode) SetMode(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()
		c = context.WithValue(c, "dev", m.Dev)
		c = context.WithValue(c, "prod", m.Prod)
		next.ServeHTTP(w, r)
	})
}
