package handler

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

type Session struct {
	manager *scs.SessionManager
}

// Sets a flash message with a given key and value.
// Honestly, this is just a convenience method because
// it doesn't do anything but set a session value.
// -- Flash messages are automatically retrieved
// an placed in the application context by the CtxFlash
// middleware for use on different parts of the page
// current flash keys:
//
//	success, error, notification,
//	emailError, passwordError, confirmError
func (hs *Session) SetFlash(c echo.Context, key, value string) {
	hs.manager.Put(c.Request().Context(), key, value)
}

// SetAuth sets an authenticated boolean and a user
// email.
func (hs Session) SetAuth(c echo.Context, email string) {
	hs.manager.Put(c.Request().Context(), "authenticated", true)
	hs.manager.Put(c.Request().Context(), "user", email)
}

// RevokeAuth removes the authenticated boolean and user string
// from the session. Basically, this logs the user out.
func (hs Session) RevokeAuth(c echo.Context) {
	hs.manager.Remove(c.Request().Context(), "authenticated")
	hs.manager.Remove(c.Request().Context(), "user")
}

// CheckAuth checks the session for "authenticated" = true
func (hs Session) CheckAuth(c echo.Context) bool {
	return hs.manager.GetBool(c.Request().Context(), "authenticated")
}

// GetUser retrieves the "user" string from session
func (hs Session) GetUser(c echo.Context) string {
	return hs.manager.GetString(c.Request().Context(), "user")
}
