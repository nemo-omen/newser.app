package handler

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
	"newser.app/server/service"
	"newser.app/shared/util"
	"newser.app/view/pages/desk"
)

type DeskHandler struct {
	session             Session
	API                 service.API                 // remote RSS feeds
	AuthService         service.AuthService         // auth logic (logout, etc)
	SubscriptionService service.SubscriptionService // subscription logic
	// NoteService										// notes logic
}

func NewDeskHandler(
	api service.API,
	ss service.SubscriptionService,
	as service.AuthService,
	sessionManager *scs.SessionManager,
) DeskHandler {
	return DeskHandler{
		session:             Session{Manager: sessionManager},
		API:                 api,
		AuthService:         as,
		SubscriptionService: ss,
	}
}

func (h DeskHandler) GetDeskIndex(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	email := h.session.GetUser(c)

	if !authed {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	u, _ := h.AuthService.GetUserByEmail(email)
	subscriptions, err := h.SubscriptionService.All(u.Id)
	if err != nil {
		// TODO: Figure out what to do when an error
		//		 is thrown here.
		fmt.Println(err.Error())
	}
	if len(subscriptions) < 1 {
		return c.Redirect(http.StatusSeeOther, "/desk/search")
	}

	return render(c, desk.Index())
}

func (h DeskHandler) GetDeskSearch(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	// email := h.session.GetUser(c)

	if !authed {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	return render(c, desk.Search([]*gofeed.Feed{}))
}

func (h DeskHandler) PostDeskSearch(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	// email := h.session.GetUser(c)

	if !authed {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	feeds := []*gofeed.Feed{}
	searchLinks := []string{}
	// isHx := c.Get("isHx").(bool)
	searchInput := c.Request().FormValue("searchurl")
	validUrl, err := util.MakeUrl(searchInput)
	if err != nil {
		h.session.SetFlash(c, "searchError", fmt.Sprintf("No feeds found at %v, try using %v.com?", searchInput, searchInput))
		return render(c, desk.Search(feeds))
	}
	isSite := h.API.CheckSite(validUrl.String())
	if !isSite {
		h.session.SetFlash(c, "searchError", fmt.Sprintf("Could not find a feed at %v", searchInput))
		return render(c, desk.Search(feeds))
	}

	docFeedLinks, _ := h.API.FindFeedLinks(validUrl.String())
	searchLinks = append(searchLinks, docFeedLinks...)

	if len(searchLinks) < 1 {
		guessedLinks, err := h.API.GuessFeedLinks(validUrl.String())
		if err != nil {
			h.session.SetFlash(c, "searchError", fmt.Sprintf("Could not fing a feed at %v", searchInput))
			return render(c, desk.Search(feeds))
		}
		searchLinks = append(searchLinks, guessedLinks...)
	}

	feedsResult, err := h.API.GetFeedsConcurrent(searchLinks)
	if err != nil {
		h.session.SetFlash(c, "searchError", ErrorFeedNotFound(searchInput))
		// TODO: find a way to send search result partials
		// and separate flash message partials
		// maybe SSE? with a flash element connected?
		// if isHx {
		// 	return render(c, desk.Search(feeds))
		// }
		return render(c, desk.Search(feeds))
	}
	feeds = append(feeds, feedsResult...)
	fmt.Println(feeds)

	// TODO: isHx => partial
	// if isHx {
	// 	return render(c, component.FeedSearchResult(feeds))
	// }
	return render(c, desk.Search(feeds))
}

func (h DeskHandler) PostDeskSubscribe(c echo.Context) error {
	authed := h.session.CheckAuth(c)
	email := h.session.GetUser(c)
	if !authed {
		h.session.SetFlash(c, "error", ErrorNotLoggedIn)
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	subscriptionUrl := c.Request().FormValue("subscriptionurl")
	feed, err := h.API.GetFeed(subscriptionUrl)
	if err != nil {
		h.session.SetFlash(c, "error", ErrorFeedNotFound(subscriptionUrl))
		// TODO: need a more suitable response here
		return render(c, desk.Index())
	}
	u, _ := h.AuthService.GetUserByEmail(email)
	sub, err := h.SubscriptionService.Subscribe(feed, int(u.Id))
	if err != nil {
		h.session.SetFlash(c, "error", fmt.Sprintf("Could not subscribe to %v", feed.Title))
	}
	fmt.Println(sub)
	return render(c, desk.Index())
}
