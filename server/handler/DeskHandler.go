package handler

import (
	"fmt"
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
	email := h.session.GetUser(c)

	if !authed {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	u, _ := h.AuthService.GetUserByEmail(email)
	subscriptions, err := h.SubscriptionService.All(u.Id)
	if err != nil {
		// TODO: Figure out what to do when an error
		//		 is thrown here.
		fmt.Println(err.Error())
	}
	if len(subscriptions) < 1 {
		return c.Redirect(http.StatusSeeOther, "/desk/search")
	}

	return render(c, desk.Index())
}

func (h DeskHandler) GetDeskSearch(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	// email := h.session.GetUser(c)

	if !authed {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	return render(c, desk.Search())
}

func (h DeskHandler) PostDeskSearch(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	// email := h.session.GetUser(c)

	if !authed {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	searchurl := c.Request().FormValue("searchurl")
	fmt.Println(searchurl)
	// TODO: Try to reimplement main branch search logic
	return render(c, desk.Search())
}
