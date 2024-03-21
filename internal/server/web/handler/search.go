package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
	"newser.app/internal/usecase/discovery"
	"newser.app/internal/usecase/session"
	"newser.app/shared/util"
	"newser.app/view/pages/app/search"
)

var (
	ErrFeedNotFound = func(resource string) string { return fmt.Sprintf("Feed not found at %v", resource) }
)

type WebSearchHandler struct {
	session session.SessionService
	// authService     auth.AuthService
	// collectionService collection.CollectionService
	// newsfeedService newsfeed.NewsfeedService
	// subscriptionService subscription.SubscriptionService
	discoveryService discovery.DiscoveryService
}

func NewWebSearchHandler(
	sessionService session.SessionService,
	// authService auth.AuthService,
	// collectionService collection.CollectionService,
	// newsfeedService newsfeed.NewsfeedService,
	// subscriptionService subscription.SubscriptionService,
	discoveryService discovery.DiscoveryService,
) *WebSearchHandler {
	return &WebSearchHandler{
		session: sessionService,
		// authService:     authService,
		// collectionService: collectionService,
		// newsfeedService: newsfeedService,
		// subscriptionService: subscriptionService,
		discoveryService: discoveryService,
	}
}

func (h *WebSearchHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/app/search", h.GetSearch)
	app.POST("/app/search", h.PostSearch)
}

func (h *WebSearchHandler) GetSearch(c echo.Context) error {
	c.Set("title", "Add a Feed")
	authed, ok := c.Get("authenticated").(bool)
	if !ok {
		authed = false
	}
	if !authed {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	util.SetPageTitle(c, h.session, "Add a Feed")
	if isHxRequest(c) {
		return render(c, search.SearchPageContent([]*gofeed.Feed{}))
	}

	return render(c, search.Search([]*gofeed.Feed{}))
}

func (h *WebSearchHandler) PostSearch(c echo.Context) error {
	// TODO: refactor so we use AppErrors and
	// renderOrRedirect when applicable
	// also, some of the logic here should
	// be moved to the discovery service
	feeds := []*gofeed.Feed{}
	searchLinks := []string{}
	searchInput := c.Request().FormValue("searchurl")
	validUrl, err := util.MakeUrl(searchInput)
	if err != nil {
		h.session.SetFlash(c, "searchError", fmt.Sprintf("No feeds found at %v, try using %v.com?", searchInput, searchInput))
		return render(c, search.Search(feeds))
	}
	isSite := h.discoveryService.IsValidSite(validUrl.String())
	if !isSite {
		h.session.SetFlash(c, "searchError", fmt.Sprintf("Could not find a feed at %v", searchInput))
		return render(c, search.Search(feeds))
	}

	docFeedLinks, _ := h.discoveryService.FindFeedLinks(validUrl.String())
	searchLinks = append(searchLinks, docFeedLinks...)

	if len(searchLinks) < 1 {
		guessedLinks, err := h.discoveryService.GuessFeedLinks(validUrl.String())
		if err != nil {
			h.session.SetFlash(c, "searchError", fmt.Sprintf("Could not find a feed at %v", searchInput))
			return render(c, search.Search(feeds))
		}
		searchLinks = append(searchLinks, guessedLinks...)
	}

	feedsResult, err := h.discoveryService.GetFeedsConcurrent(searchLinks)
	if err != nil {
		h.session.SetFlash(c, "searchError", ErrFeedNotFound(searchInput))
		return render(c, search.Search(feeds))
	}
	feeds = append(feeds, feedsResult...)
	return render(c, search.Search(feeds))
}
