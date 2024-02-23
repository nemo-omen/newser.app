package handler

import (
	"current/view/app"
	"fmt"

	"github.com/labstack/echo/v5"
)

type AppHandler struct{}

func (h AppHandler) HandleGetIndex(c echo.Context) error {
	fmt.Println(c)
	return render(c, app.Index())
}
