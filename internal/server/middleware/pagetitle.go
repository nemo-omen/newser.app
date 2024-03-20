package middleware

import (
	"fmt"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func PageTitle(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			pageTitle := getPageTitle(c, sm)
			fmt.Println("middleware pageTitle: ", pageTitle)
			c.Set("title", pageTitle)
			return next(c)
		}
	}
}

// getFlash is a convenience function which pops
// a flash message with a given key and returns its value
func getPageTitle(c echo.Context, sm *scs.SessionManager) string {
	return sm.PopString(c.Request().Context(), "title")
}
