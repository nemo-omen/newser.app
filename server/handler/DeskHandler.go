package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"newser.app/model"
	"newser.app/server/service"
	"newser.app/shared/util"
	"newser.app/view/component"
	"newser.app/view/pages/desk"
)

type DeskHandler struct {
	session             Session
	api                 service.API                 // remote RSS feeds
	authService         service.AuthService         // auth logic (logout, etc)
	subscriptionService service.SubscriptionService // subscription logic
	newsfeedService     service.NewsfeedService
	collectionService   service.CollectionService
	// NoteService										// notes logic
}

func NewDeskHandler(
	api service.API,
	ss service.SubscriptionService,
	as service.AuthService,
	ns service.NewsfeedService,
	cs service.CollectionService,
	sessionManager *scs.SessionManager,
) DeskHandler {
	return DeskHandler{
		session:             Session{manager: sessionManager},
		api:                 api,
		authService:         as,
		subscriptionService: ss,
		newsfeedService:     ns,
		collectionService:   cs,
	}
}

func (h DeskHandler) GetDeskIndex(c echo.Context) error {
	eml := h.session.GetUser(c)
	user, err := h.authService.GetUserByEmail(eml)
	isHx := c.Get("isHx")
	// ref := c.Request().Referer()

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
	if isHx != nil {
		if isHx.(bool) {
			return render(c, desk.IndexPageContent(storedSubscriptionArticles))
		}
	}
	return render(c, desk.Index(storedSubscriptionArticles))
}

func (h DeskHandler) GetDeskSearch(c echo.Context) error {
	isHx := c.Get("isHx")
	c.Set("title", "Add a Newsfeed")
	if isHx != nil {
		if isHx.(bool) {
			return render(c, desk.SearchPageContent([]*gofeed.Feed{}))
		}
	}
	return render(c, desk.Search([]*gofeed.Feed{}))
}

func (h DeskHandler) GetDeskArticle(c echo.Context) error {
	stringId := c.Param("articleid")
	id, err := strconv.ParseInt(stringId, 10, 64)
	isHx := c.Get("isHx")

	if err != nil {
		h.session.SetFlash(c, "error", "Error retrieving article.")
		return render(c, desk.Article((&model.Article{Title: "Oops!"})))
	}

	article, err := h.newsfeedService.GetArticleById(id)
	if err != nil {
		h.session.SetFlash(c, "error", "Error retrieving article.")
		return render(c, desk.Article((&model.Article{Title: "Oops!"})))
	}
	c.Set("title", article.FeedTitle)
	if isHx != nil {
		if isHx.(bool) {
			return render(c, desk.ArticlePageContent(article))
		}
	}
	return render(c, desk.Article(article))
}

func (h DeskHandler) GetDeskNewsfeed(c echo.Context) error {
	idStr := c.Param("feedid")
	id, err := strconv.ParseInt(idStr, 10, 64)
	isHx := c.Get("isHx")

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
	if isHx != nil {
		if isHx.(bool) {
			return render(c, desk.NewsfeedPageContent(feed))
		}
	}
	return render(c, desk.Newsfeed(feed))
}

func (h *DeskHandler) GetDeskCollection(c echo.Context) error {
	collectionName := c.Param("collectionname")
	email := h.session.GetUser(c)
	isHx := c.Get("isHx")
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		h.session.SetFlash(c, "error", "You need to be logged in")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	collectionArticles, err := h.collectionService.GetArticlesByCollectionByName(collectionName, user.Id)
	upper := cases.Title(language.AmericanEnglish).String(collectionName)
	c.Set("title", upper)
	if err != nil {
		h.session.SetFlash(c, "error", fmt.Sprintf("Error getting %s articles", upper))
	}

	if isHx != nil {
		if isHx.(bool) {
			return render(c, desk.IndexPageContent(collectionArticles))
		}
	}
	return render(c, desk.Index(collectionArticles))
}

func (h *DeskHandler) GetDeskUnreadCount(c echo.Context) error {
	return render(c, component.MainFeedLinks())
}

func (h DeskHandler) PostDeskSearch(c echo.Context) error {
	feeds := []*gofeed.Feed{}
	searchLinks := []string{}
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

func (h DeskHandler) PostDeskAddToRead(c echo.Context) error {
	email := h.session.GetUser(c)
	user, err := h.authService.GetUserByEmail(email)
	isHx := c.Get("isHx")
	if err != nil {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	ref := c.Request().Referer()
	fmt.Println("referrer: ", ref)

	idStr := c.Request().FormValue("articleid")
	fmt.Println("idStr", idStr)
	if idStr == "" {
		h.session.SetFlash(c, "error", "Could not mark as read")
		return c.Redirect(http.StatusSeeOther, ref)
	}

	aId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.session.SetFlash(c, "error", "Could not mark as read")
		return c.Redirect(http.StatusSeeOther, ref)
	}

	fmt.Println("user: ", user.Id)
	fmt.Println("article: ", aId)
	err = h.collectionService.AddArticleToRead(aId, user.Id)
	if err != nil {
		h.session.SetFlash(c, "error", "Could not mark as read")
		return c.Redirect(http.StatusSeeOther, ref)
	}

	if isHx != nil {
		article, err := h.newsfeedService.GetArticleById(aId)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, ref)
		}

		if isHx.(bool) {
			c.Response().Header().Add("HX-Trigger", "updateUnreadCount")
			return render(c, component.ArticleCard(article))
		}
	}

	return c.Redirect(http.StatusSeeOther, ref)
}

func (h DeskHandler) PostDeskAddToUnread(c echo.Context) error {
	email := h.session.GetUser(c)
	user, err := h.authService.GetUserByEmail(email)
	isHx := c.Get("isHx")
	if err != nil {
		h.session.SetFlash(c, "error", "You need to log in.")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	ref := c.Request().Referer()
	fmt.Println("referrer: ", ref)

	idStr := c.Request().FormValue("articleid")
	fmt.Println("idStr", idStr)
	if idStr == "" {
		h.session.SetFlash(c, "error", "Could not mark as unread")
		return c.Redirect(http.StatusSeeOther, ref)
	}

	aId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.session.SetFlash(c, "error", "Could not mark as unread")
		return c.Redirect(http.StatusSeeOther, ref)
	}

	fmt.Println("user: ", user.Id)
	fmt.Println("article: ", aId)
	err = h.collectionService.RemoveArticleFromRead(aId, user.Id)
	if err != nil {
		h.session.SetFlash(c, "error", "Could not mark as unread")
		return c.Redirect(http.StatusSeeOther, ref)
	}

	if isHx != nil {
		article, err := h.newsfeedService.GetArticleById(aId)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, ref)
		}

		if isHx.(bool) {
			c.Response().Header().Add("HX-Trigger", "updateUnreadCount")
			return render(c, component.ArticleCard(article))
		}
	}

	return c.Redirect(http.StatusSeeOther, ref)
}

func (h DeskHandler) DeskPostCardCollapsed(c echo.Context) error {
	ref := c.Request().Referer()
	idInput := c.Request().FormValue("articleid")
	collapseInput := c.Request().FormValue("shouldcollapse")
	isHx := c.Get("isHx")

	id, err := strconv.ParseInt(idInput, 10, 64)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, ref)
	}

	if collapseInput == "true" {
		h.session.PutCollapsedCard(c, id)
	} else {
		if h.session.HasCollapsedCard(c, id) {
			h.session.RemoveCollapsedCard(c, id)
		}
	}

	if isHx != nil {
		article, err := h.newsfeedService.GetArticleById(id)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, ref)
		}

		if isHx.(bool) {
			return render(c, component.ArticleCard(article))
		}
	}
	fmt.Println("why are you here?")

	return c.Redirect(http.StatusSeeOther, ref)
}
