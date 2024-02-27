package handler

import (
	"github.com/labstack/echo/v4"
	"newser.app/server/service"
	"newser.app/view/pages/desk"
)

type DeskHandler struct {
	API         service.API         // remote RSS feeds
	AuthService service.AuthService // auth logic (logout, etc)
	// SubscriptionService service.SubscriptionService // subscription logic
	// NoteService										// notes logic
}

func NewDeskHandler(api service.API, ss service.SubscriptionService, as service.AuthService) DeskHandler {
	return DeskHandler{
		API:         api,
		AuthService: as,
		// SubscriptionService: ss,
	}
}

func (h DeskHandler) GetDeskIndex(c echo.Context) error {
	getFlashes(c)
	return render(c, desk.Index())
}
