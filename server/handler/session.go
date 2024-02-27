package handler

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var sessionStore = sessions.NewCookieStore([]byte("super-secret-session-key"))

func SetAuthedSession(c echo.Context) {
	session, _ := sessionStore.Get(c.Request(), "user_auth")
	session.Options.MaxAge = int((7 * 24) * time.Hour)
	session.Options.HttpOnly = true
	session.Options.SameSite = http.SameSiteStrictMode
	session.Values["authenticated"] = true
	session.Save(c.Request(), c.Response())
}

func RevokeAuthedSession(c echo.Context) {
	session, _ := sessionStore.Get(c.Request(), "user_auth")
	session.Values["authenticated"] = false
	session.Save(c.Request(), c.Response())
}

func CheckAuthedSession(c echo.Context) bool {
	session, _ := sessionStore.Get(c.Request(), "user_auth")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}
