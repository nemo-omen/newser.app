package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func HashPassword(p string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(p), 8)
	return string(hashed)
}

func PasswordMatches(p, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p), []byte(h))
	return err == nil
}

func getFlashes(c echo.Context) {
	_ = GetFlash(c, "errorFlash")
	_ = GetFlash(c, "successFlash")
	_ = GetFlash(c, "notificationFlash")
}
