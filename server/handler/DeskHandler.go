package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
	"newser.app/model"
	"newser.app/server/service"
	"newser.app/shared/util"
	"newser.app/view/pages/desk"
)

type DeskHandler struct {
	session             Session
	api                 service.API                 // remote RSS feeds
	authService         service.AuthService         // auth logic (logout, etc)
	subscriptionService service.SubscriptionService // subscription logic
	newsfeedService     service.NewsfeedService
	// NoteService										// notes logic
}

func NewDeskHandler(
	api service.API,
	ss service.SubscriptionService,
	as service.AuthService,
	ns service.NewsfeedService,
	sessionManager *scs.SessionManager,
) DeskHandler {
	return DeskHandler{
		session:             Session{manager: sessionManager},
		api:                 api,
		authService:         as,
		subscriptionService: ss,
		newsfeedService:     ns,
	}
}

func (h DeskHandler) GetDeskIndex(c echo.Context) error {
	eml := h.session.GetUser(c)
	user, err := h.authService.GetUserByEmail(eml)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	storedSubscriptionArticles, err := h.subscriptionService.GetArticles(user.Id)

	if err != nil {
		h.session.SetFlash(c, "error", "Error getting your feeds")
		fmt.Println("error getting stored subscription articles: ", err.Error())
	}

	if len(storedSubscriptionArticles) < 1 {
		return c.Redirect(http.StatusSeeOther, "/desk/search")
	}

	c.Set("title", "Latest Articles")
	return render(c, desk.Index(storedSubscriptionArticles))
}

func (h DeskHandler) GetDeskSearch(c echo.Context) error {
	c.Set("title", "Add a Newsfeed")
	return render(c, desk.Search([]*gofeed.Feed{}))
}

func (h DeskHandler) GetDeskArticle(c echo.Context) error {
	stringId := c.Param("articleid")
	id, err := strconv.ParseInt(stringId, 10, 64)

	if err != nil {
		h.session.SetFlash(c, "error", "Error retrieving article.")
		return render(c, desk.Article((&model.Article{Title: "Oops!"})))
	}

	article, err := h.newsfeedService.GetArticleById(id)
	if err != nil {
		h.session.SetFlash(c, "error", "Error retrieving article.")
		return render(c, desk.Article((&model.Article{Title: "Oops!"})))
	}
	c.Set("title", article.Title)
	return render(c, desk.Article(article))
}

func (h DeskHandler) GetDeskNewsfeed(c echo.Context) error {
	idStr := c.Param("feedid")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.session.SetFlash(c, "error", "Error getting newsfeed")
		return render(c, desk.Newsfeed(&model.Newsfeed{}))
	}

	feed, err := h.newsfeedService.GetNewsfeed(id)
	if err != nil {
		h.session.SetFlash(c, "error", "Error getting newsfeed.")
		return render(c, desk.Newsfeed(&model.Newsfeed{}))
	}
	c.Set("title", feed.Title)
	return render(c, desk.Newsfeed(feed))
}

func (h DeskHandler) PostDeskSearch(c echo.Context) error {
	feeds := []*gofeed.Feed{}
	searchLinks := []string{}
	// isHx := c.Get("isHx").(bool)
	searchInput := c.Request().FormValue("searchurl")
	validUrl, err := util.MakeUrl(searchInput)
	if err != nil {
		h.session.SetFlash(c, "searchError", fmt.Sprintf("No feeds found at %v, try using %v.com?", searchInput, searchInput))
		return render(c, desk.Search(feeds))
	}
	isSite := h.api.CheckSite(validUrl.String())
	if !isSite {
		h.session.SetFlash(c, "searchError", fmt.Sprintf("Could not find a feed at %v", searchInput))
		return render(c, desk.Search(feeds))
	}

	docFeedLinks, _ := h.api.FindFeedLinks(validUrl.String())
	searchLinks = append(searchLinks, docFeedLinks...)

	if len(searchLinks) < 1 {
		guessedLinks, err := h.api.GuessFeedLinks(validUrl.String())
		if err != nil {
			h.session.SetFlash(c, "searchError", fmt.Sprintf("Could not find a feed at %v", searchInput))
			return render(c, desk.Search(feeds))
		}
		searchLinks = append(searchLinks, guessedLinks...)
	}

	feedsResult, err := h.api.GetFeedsConcurrent(searchLinks)
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
	// fmt.Println(feeds)

	// TODO: isHx => partial
	// if isHx {
	// 	return render(c, component.FeedSearchResult(feeds))
	// }
	return render(c, desk.Search(feeds))
}

func (h DeskHandler) PostDeskSubscribe(c echo.Context) error {
	email := h.session.GetUser(c)
	subscriptionUrl := c.Request().FormValue("subscriptionurl")
	fmt.Println("subscriptionurl: ", subscriptionUrl)
	feed, err := h.api.GetFeed(subscriptionUrl)
	if err != nil {
		fmt.Println("feed error: ", err.Error())
		h.session.SetFlash(c, "error", ErrorFeedNotFound(subscriptionUrl))
		// TODO: need a more suitable response here
		return c.Redirect(http.StatusSeeOther, "/desk/")
	}
	u, _ := h.authService.GetUserByEmail(email)
	newsfeed, err := h.subscriptionService.Subscribe(feed, u.Id)
	if err != nil {
		h.session.SetFlash(c, "error", fmt.Sprintf("Could not subscribe to %v", feed.Title))
	}
	// fmt.Println(sub)
	// return render(c, desk.Index())
	h.session.SetFlash(c, "success", fmt.Sprintf("success subscribing to %v", newsfeed.Title))
	return c.Redirect(http.StatusSeeOther, "/desk/feeds/"+strconv.FormatInt(newsfeed.ID, 10))
}
