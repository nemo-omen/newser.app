package handler

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"newser.app/server/service"
	"newser.app/view/pages/desk"
)

type DeskHandler struct {
	session             Session
	API                 service.API                 // remote RSS feeds
	AuthService         service.AuthService         // auth logic (logout, etc)
	SubscriptionService service.SubscriptionService // subscription logic
	// NoteService										// notes logic
}

func NewDeskHandler(
	api service.API,
	ss service.SubscriptionService,
	as service.AuthService,
	sessionManager *scs.SessionManager,
) DeskHandler {
	return DeskHandler{
		session:             Session{Manager: sessionManager},
		API:                 api,
		AuthService:         as,
		SubscriptionService: ss,
	}
}

func (h DeskHandler) GetDeskIndex(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	if !authed {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	// subscriptions
	return render(c, desk.Index())
}
