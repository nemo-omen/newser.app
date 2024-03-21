package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func ViewPreference(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("view", getView(c, sm))
			return next(c)
		}
	}
}

func getView(c echo.Context, sm *scs.SessionManager) string {
	view := sm.GetString(c.Request().Context(), "view")
	if view == "" {
		return "card"
	}
	return view
}
