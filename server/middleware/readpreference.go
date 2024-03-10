package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func ReadPreference(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("viewRead", getRead(c, sm))
			return next(c)
		}
	}
}

func getRead(c echo.Context, sm *scs.SessionManager) bool {
	return sm.GetBool(c.Request().Context(), "viewRead")
}
