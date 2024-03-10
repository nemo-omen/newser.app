package handler

import (
	"slices"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"newser.app/shared/util"
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

func (hs Session) SetCollapsedCards(c echo.Context, collapsedCards []int64) {
	hs.manager.Put(c.Request().Context(), "collapsedcards", collapsedCards)
	hs.manager.Commit(c.Request().Context())
	c.Set("collapsedCards", collapsedCards)
}

func (hs Session) GetCollapsedCards(c echo.Context) []int64 {
	collapsedSession := hs.manager.Get(c.Request().Context(), "collapsedcards")
	collapsedCards := []int64{}
	if collapsedSession != nil {
		collapsedCards = collapsedSession.([]int64)
	}
	return collapsedCards
}

func (hs Session) PutCollapsedCard(c echo.Context, card int64) {
	collapsedCards := hs.GetCollapsedCards(c)
	if !slices.Contains(collapsedCards, card) {
		collapsedCards = append(collapsedCards, card)
	}
	hs.SetCollapsedCards(c, collapsedCards)
}

func (hs Session) RemoveCollapsedCard(c echo.Context, card int64) {
	collapsedCards := hs.GetCollapsedCards(c)
	if len(collapsedCards) < 1 {
		return
	}
	filtered := util.Filter(collapsedCards, func(cardId int64) bool {
		return cardId != card
	})
	hs.SetCollapsedCards(c, filtered)
}

func (hs Session) HasCollapsedCard(c echo.Context, card int64) bool {
	collapsedCards := hs.GetCollapsedCards(c)
	if len(collapsedCards) < 1 {
		return false
	}
	return slices.Contains(collapsedCards, card)
}

func (hs *Session) SetTitle(c echo.Context, title string) {
	hs.manager.Put(c.Request().Context(), "pagetitle", title)
	hs.manager.Commit(c.Request().Context())
}

func (hs *Session) SetView(c echo.Context, view string) {
	hs.manager.Put(c.Request().Context(), "view", view)
	hs.manager.Commit(c.Request().Context())
	c.Set("view", view)
}

func (hs *Session) GetView(c echo.Context) string {
	view := hs.manager.GetString(c.Request().Context(), "view")
	if view == "" {
		return "card"
	}
	return view
}
