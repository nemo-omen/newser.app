package util

import (
	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/session"
)

func SetPageTitle(c echo.Context, s session.SessionService, title string) {
	s.SetTitle(c, title)
	c.Response().Header().Set("HX-Trigger", "pagetitle")
}
