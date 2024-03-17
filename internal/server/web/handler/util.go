package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func isHxRequest(c echo.Context) bool {
	return c.Get("isHx") != nil && c.Get("isHx").(bool)
}
