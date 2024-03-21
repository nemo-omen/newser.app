package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

// "read" or "unread"
func ReadPreference(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("view", getRead(c, sm))
			return next(c)
		}
	}
}

// "read" or "unread"
func getRead(c echo.Context, sm *scs.SessionManager) bool {
	return sm.GetBool(c.Request().Context(), "view")
}
