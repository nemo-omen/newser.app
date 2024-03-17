package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/session"
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
	authed := h.session.GetAuth(c)
	fmt.Println("GET /", authed)
	return render(c, home.Index())
}
