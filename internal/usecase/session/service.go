package session

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

type SessionService struct {
	sessionManager *scs.SessionManager
}

func NewSessionService(sessionManager *scs.SessionManager) SessionService {
	return SessionService{
		sessionManager: sessionManager,
	}
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
func (s *SessionService) SetFlash(c echo.Context, key, value string) {
	s.sessionManager.Put(c.Request().Context(), key, value)
}

func (s *SessionService) GetFlash(c echo.Context, key string) string {
	if !s.sessionManager.Exists(c.Request().Context(), key) {
		return ""
	}
	return s.sessionManager.PopString(c.Request().Context(), key)
}

func (s *SessionService) SetUser(c echo.Context, user string) {
	s.sessionManager.Put(c.Request().Context(), "user", user)
}

func (s *SessionService) GetUser(c echo.Context) string {
	if !s.sessionManager.Exists(c.Request().Context(), "user") {
		return ""
	}
	return s.sessionManager.GetString(c.Request().Context(), "user")
}

func (s *SessionService) SetAuth(c echo.Context, user string) {
	s.SetUser(c, user)
	s.sessionManager.Put(c.Request().Context(), "authenticated", true)
}

func (s *SessionService) GetAuth(c echo.Context) bool {
	if !s.sessionManager.Exists(c.Request().Context(), "authenticated") {
		return false
	}
	return s.sessionManager.GetBool(c.Request().Context(), "authenticated")
}

func (s *SessionService) RevokeAuth(c echo.Context) {
	s.sessionManager.Remove(c.Request().Context(), "user")
	s.sessionManager.Put(c.Request().Context(), "authenticated", false)
}

func (s *SessionService) SetTitle(c echo.Context, title string) {
	// fmt.Println("setting title: ", title)
	s.sessionManager.Put(c.Request().Context(), "title", title)
}

func (s *SessionService) GetTitle(c echo.Context) string {
	if !s.sessionManager.Exists(c.Request().Context(), "title") {
		return "Newser"
	}
	return s.sessionManager.GetString(c.Request().Context(), "title")
}

// "expanded" or "collapsed"
func (s *SessionService) SetLayout(c echo.Context, layout string) {
	s.sessionManager.Put(c.Request().Context(), "layout", layout)
}

// "expanded" or "collapsed"
func (s *SessionService) GetLayout(c echo.Context) string {
	if !s.sessionManager.Exists(c.Request().Context(), "layout") {
		return "expanded"
	}
	return s.sessionManager.GetString(c.Request().Context(), "layout")
}

// "read" or "unread"
func (s *SessionService) SetView(c echo.Context, view string) {
	s.sessionManager.Put(c.Request().Context(), "view", view)
}

// "read" or "unread"
// "unread" is default
func (s *SessionService) GetView(c echo.Context) string {
	if !s.sessionManager.Exists(c.Request().Context(), "view") {
		return "unread"
	}
	return s.sessionManager.GetString(c.Request().Context(), "view")
}
