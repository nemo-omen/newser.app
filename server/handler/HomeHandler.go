package handler

import (
	"github.com/labstack/echo/v4"
	"newser.app/view/pages/home"
)

type HomeHandler struct{}

// create route handlers here
func (h HomeHandler) Home(c echo.Context) error {
	return render(c, home.Index())
}
