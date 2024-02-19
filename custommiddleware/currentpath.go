package custommiddleware

import (
	"github.com/labstack/echo/v4"
)

func CurrentPath(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("currentPath", c.Request().URL.Path)
		return next(c)
	}
}
