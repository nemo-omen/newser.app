package middleware

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func CtxCardState(sm *scs.SessionManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("collapsedCards", getCollapsedCards(c, sm))
			return next(c)
		}
	}
}

func getCollapsedCards(c echo.Context, sm *scs.SessionManager) []int64 {
	collapsedSession := sm.Get(c.Request().Context(), "collapsedcards")
	collapsedCards := []int64{}
	if collapsedSession != nil {
		collapsedCards = collapsedSession.([]int64)
	}
	return collapsedCards
}
