package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

// AuthContext middleware uses a given sessionManager
// to retrieve auth-related data from the current session
// and adds that data to context, making it available
// to handlers and views
func AuthContext(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isAuthed := sm.GetBool(c.Request().Context(), "authenticated")
			user := sm.GetString(c.Request().Context(), "user")
			c.Set("authenticated", isAuthed)
			c.Set("user", user)
			return next(c)
		}
	}
}
