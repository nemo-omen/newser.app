package middleware

import (
	"github.com/labstack/echo/v4"
	"newser.app/server/handler"
)

func CtxFlash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_ = handler.GetFlash(c, "errorFlash")
		_ = handler.GetFlash(c, "successFlash")
		_ = handler.GetFlash(c, "notificationFlash")
		_ = handler.GetFlash(c, "errorFlash")
		_ = handler.GetFlash(c, "successFlash")
		_ = handler.GetFlash(c, "notificationFlash")
		_ = handler.GetFlash(c, "emailError")
		_ = handler.GetFlash(c, "passwordError")
		_ = handler.GetFlash(c, "confirmError")
		return next(c)
	}
}
