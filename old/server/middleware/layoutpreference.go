package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

// "expanded" or "collapsed
func LayoutPreference(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("layout", getLayout(c, sm))
			return next(c)
		}
	}
}

// "expanded" or "collapsed"
// if the session value is empty, it will default to "expanded"
func getLayout(c echo.Context, sm *scs.SessionManager) string {
	view := sm.GetString(c.Request().Context(), "layout")
	if view == "" {
		return "expanded"
	}
	return view
}
