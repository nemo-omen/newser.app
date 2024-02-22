package handler

import (
	"current/view/home"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleGetIndex(c echo.Context) error {
	user := c.Get("user")
	fmt.Println("user: ", user)
	if user != nil {
		return c.Redirect(http.StatusFound, "/app/")
	}
	return render(c, home.Index())
}
