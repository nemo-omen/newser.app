package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func Auth(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authed := sm.GetBool(c.Request().Context(), "authenticated")
			if !authed {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			}
			return next(c)
		}
	}
}
