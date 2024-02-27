package middleware

import (
	"github.com/labstack/echo/v4"
	"newser.app/server/handler"
)

func AuthContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuthed := handler.CheckAuthedSession(c)
		c.Set("authenticated", isAuthed)
		return next(c)
	}
}
