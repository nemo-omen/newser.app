package handler

import (
	"net/http"

	"newser.app/internal/service"
)

type AppHandler struct {
	API                 service.API                 // remote RSS feeds
	AuthService         service.AuthService         // auth logic (logout, etc)
	SubscriptionService service.SubscriptionService // subscription logic
	// NoteService										// notes logic
}

func NewAppHandler(dsn string) AppHandler {
	return AppHandler{
		API:                 *service.NewAPI(&http.Client{}),
		AuthService:         service.NewAuthService(dsn),
		SubscriptionService: service.NewSubscriptionService(dsn),
	}
}
