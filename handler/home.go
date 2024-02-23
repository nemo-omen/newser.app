package handler

import (
	"current/view/home"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleGetIndex(c echo.Context) error {
	sess, _ := session.Get("session", c)
	fmt.Println(sess)
	user := c.Get("user")
	fmt.Println("user: ", user)
	if user != nil {
		return c.Redirect(http.StatusFound, "/app/")
	}
	return render(c, home.Index())
}
