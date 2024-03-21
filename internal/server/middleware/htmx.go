package middleware

import (
	"github.com/labstack/echo/v4"
)

func HTMX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hxHeader := c.Request().Header.Get("HX-Request")
		if hxHeader != "" {
			c.Set("isHx", true)
		} else {
			c.Set("isHx", false)
		}
		return next(c)
	}
}
