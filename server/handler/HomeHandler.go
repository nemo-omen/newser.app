package handler

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"newser.app/view/pages/home"
)

type HomeHandler struct {
	session Session
}

func NewHomeHandler(sessionManager *scs.SessionManager) HomeHandler {
	return HomeHandler{
		session: Session{Manager: sessionManager},
	}
}

func (h HomeHandler) Home(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	if authed {
		return c.Redirect(http.StatusSeeOther, "/desk/")
	}
	return render(c, home.Index())
}
