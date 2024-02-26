package handler

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var store = sessions.NewCookieStore([]byte("super-secret-flash-key"))

func SetFlash(c echo.Context, key, value string) {
	session, err := store.Get(c.Request(), "flash-session")
	if err != nil {
		fmt.Printf("error setting flash: %s", err.Error())
	}
	session.AddFlash(value, key)
	session.Save(c.Request(), c.Response())
}

func GetFlash(c echo.Context, key string) string {
	session, err := store.Get(c.Request(), "flash-session")
	if err != nil {
		fmt.Printf("error getting flash message: %s\n", err.Error())
		return ""
	}

	fm := session.Flashes(key)
	if fm == nil {
		return ""
	}
	session.Save(c.Request(), c.Response())
	c.Set(key, fm[0])
	return fm[0].(string)
}
