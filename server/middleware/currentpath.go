package middleware

import (
	"github.com/labstack/echo/v4"
)

// SetCurrentPath sets a request context value
// to the path of the current request URL. This
// is used to pass the path to the templ Header component
// to avoid prop drilling.
func SetCurrentPath(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("currentPath", c.Request().URL.Path)
		return next(c)
	}
}
