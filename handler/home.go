package handler

import (
	"current/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleGetIndex(c echo.Context) error {
	return render(c, home.Index())
}
