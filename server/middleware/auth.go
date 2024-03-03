package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func Auth(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isAuthed := sm.GetBool(c.Request().Context(), "authenticated")
			if !isAuthed {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			} else {
				return next(c)
			}
		}
	}
}
