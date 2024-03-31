package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

// "read" or "unread"
func VewPreference(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("view", getView(c, sm))
			return next(c)
		}
	}
}

// "read" or "unread"
func getView(c echo.Context, sm *scs.SessionManager) string {
	return sm.GetString(c.Request().Context(), "view")
}
