package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

// CtxFlash pops flash messages from the session
// store and adds them to context for easy display
// in Templ components
func CtxFlash(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("errorFlash", getFlash(c, sm, "errorFlash"))
			c.Set("successFlash", getFlash(c, sm, "successFlash"))
			c.Set("notificationFlash", getFlash(c, sm, "notificationFlash"))
			c.Set("emailError", getFlash(c, sm, "emailError"))
			c.Set("passwordError", getFlash(c, sm, "passwordError"))
			c.Set("confirmError", getFlash(c, sm, "confirmError"))
			c.Set("searchError", getFlash(c, sm, "searchError"))
			return next(c)
		}
	}
}

// getFlash is a convenience function which pops
// a flash message with a given key and returns its value
func getFlash(c echo.Context, sm *scs.SessionManager, key string) string {
	return sm.PopString(c.Request().Context(), key)
}
