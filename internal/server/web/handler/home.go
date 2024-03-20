package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/session"
	"newser.app/shared/util"
	"newser.app/view/pages/home"
)

type WebHomeHandler struct {
	session session.SessionService
}

func NewWebHomeHandler(session session.SessionService) *WebHomeHandler {
	return &WebHomeHandler{
		session: session,
	}
}

func (h *WebHomeHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/", h.Home)
}

func (h *WebHomeHandler) Home(c echo.Context) error {
	fmt.Println("GET /")
	authed, ok := c.Get("authenticated").(bool)
	if !ok {
		authed = false
	}
	if authed {
		return c.Redirect(http.StatusSeeOther, "/app")
	}

	util.SetPageTitle(c, h.session, "Newser")
	return render(c, home.Index())
}
