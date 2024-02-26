package handler

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const (
	session_name              string = "fmessages"
	session_flashmessages_key string = "flashmessages-key"
)

func getCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(session_flashmessages_key))
}

func setFlashMessages(c echo.Context, kind, value string) {
	session, _ := getCookieStore().Get(c.Request(), session_name)
	session.AddFlash(value, kind)
	c.Set(kind, value)
	session.Save(c.Request(), c.Response())
}

func getFlashMessages(c echo.Context, kind string) []string {
	session, _ := getCookieStore().Get(c.Request(), session_name)

	fm := session.Flashes(kind)

	if len(fm) > 0 {
		session.Save(c.Request(), c.Response())

		var flashes []string
		for _, fl := range fm {
			flashes = append(flashes, fl.(string))
		}

		return flashes
	}
	return nil
}
