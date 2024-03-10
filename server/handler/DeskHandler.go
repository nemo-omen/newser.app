package handler

import (
	"fmt"
	"log"
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

func isHxRequest(c echo.Context) bool {
	return c.Get("isHx") != nil && c.Get("isHx").(bool)
}

func handleErrorFlash(c echo.Context, s Session, err string) {
	s.SetFlash(c, "error", err)
}

func setPageTitle(c echo.Context, s Session, title string) {
	s.SetTitle(c, title)
}

func hxTriggerTitleHeader(c echo.Context) echo.Context {
	c.Response().Header().Add("Hx-Trigger", "updatePageTitle")
	return c
}

func (h DeskHandler) GetDeskIndex(c echo.Context) error {
	userEmail := h.session.GetUser(c)
	user, err := h.authService.GetUserByEmail(userEmail)

	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	storedArticles, err := h.subscriptionService.GetArticles(user.Id)

	if err != nil {
		handleErrorFlash(c, h.session, "Error getting your feeds")
		log.Println("error getting stored subscription articles: ", err)
	}

	if len(storedArticles) < 1 {
		return c.Redirect(http.StatusSeeOther, "/desk/search")
	}

	setPageTitle(c, h.session, "Latest Articles")
	if isHxRequest(c) {
		// connect SSE & send updated articles after initial response
		c = hxTriggerTitleHeader(c)
		return render(c, desk.IndexPageContent(storedArticles))
	}
	// not hxRequest?
	// - Update remote feeds
	// - Wait for persistence, etc
	// - merge new articles
	// - send full list

	return render(c, desk.Index(storedArticles))
}

func (h DeskHandler) GetDeskSearch(c echo.Context) error {
	setPageTitle(c, h.session, "Add a Newsfeed")
	if isHxRequest(c) {
		c = hxTriggerTitleHeader(c)
		return render(c, desk.SearchPageContent([]*gofeed.Feed{}))
	}

	return render(c, desk.Search([]*gofeed.Feed{}))
}

func (h DeskHandler) GetDeskArticle(c echo.Context) error {
	stringId := c.Param("articleid")
	id, err := strconv.ParseInt(stringId, 10, 64)

	if err != nil {
		handleErrorFlash(c, h.session, "Error retrieving article.")
		return render(c, desk.Article((&model.Article{Title: "Oops!"})))
	}

	article, err := h.newsfeedService.GetArticleById(id)
	if err != nil {
		handleErrorFlash(c, h.session, "Error retrieving article.")
		return render(c, desk.Article((&model.Article{Title: "Oops!"})))
	}

	setPageTitle(c, h.session, article.FeedTitle)
	if isHxRequest(c) {
		c = hxTriggerTitleHeader(c)
		return render(c, desk.ArticlePageContent(article))
	}
	return render(c, desk.Article(article))
}

func (h DeskHandler) GetDeskNewsfeed(c echo.Context) error {
	idStr := c.Param("feedid")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		handleErrorFlash(c, h.session, "Error getting newsfeed")
		return render(c, desk.Newsfeed(&model.Newsfeed{}))
	}

	feed, err := h.newsfeedService.GetNewsfeed(id)
	if err != nil {
		handleErrorFlash(c, h.session, "Error getting newsfeed.")
		return render(c, desk.Newsfeed(&model.Newsfeed{}))
	}

	setPageTitle(c, h.session, feed.Title)
	if isHxRequest(c) {
		c = hxTriggerTitleHeader(c)
		return render(c, desk.NewsfeedPageContent(feed))
	}
	return render(c, desk.Newsfeed(feed))
}

func (h *DeskHandler) GetDeskCollection(c echo.Context) error {
	collectionName := c.Param("collectionname")
	email := h.session.GetUser(c)
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		h.session.SetFlash(c, "error", "You need to be logged in")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	collectionArticles, err := h.collectionService.GetArticlesByCollectionByName(collectionName, user.Id)
	upper := cases.Title(language.AmericanEnglish).String(collectionName)
	setPageTitle(c, h.session, upper)
	if err != nil {
		handleErrorFlash(c, h.session, fmt.Sprintf("Error getting %s articles", upper))
	}

	if isHxRequest(c) {
		c = hxTriggerTitleHeader(c)
		return render(c, desk.IndexPageContent(collectionArticles))
	}
	return render(c, desk.Index(collectionArticles))
}

func (h DeskHandler) GetDeskNotes(c echo.Context) error {
	h.session.SetTitle(c, "Your Notes")
	if isHxRequest(c) {
		c = hxTriggerTitleHeader(c)
		return render(c, desk.NotesPageContent())
	}
	return render(c, desk.Notes())
}

func (h *DeskHandler) GetDeskUnreadCount(c echo.Context) error {
	return render(c, component.MainFeedLinks())
}

func (h *DeskHandler) GetDeskPageTitle(c echo.Context) error {
	title := c.Get("title")
	if title == nil {
		return c.String(http.StatusOK, "")
	}
	return c.String(http.StatusOK, title.(string))
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
		return render(c, desk.Search(feeds))
	}
	feeds = append(feeds, feedsResult...)
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
	if idStr == "" {
		h.session.SetFlash(c, "error", "Could not mark as read")
		return c.Redirect(http.StatusSeeOther, ref)
	}

	aId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.session.SetFlash(c, "error", "Could not mark as read")
		return c.Redirect(http.StatusSeeOther, ref)
	}

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

		responseType := c.Request().FormValue("responseType")

		if isHx.(bool) {
			c.Response().Header().Add("HX-Trigger", "updateUnreadCount")
			if responseType == "card" {
				return render(c, component.ArticleCard(article))
			} else {
				return render(c, component.ArticleCondensed(article))
			}
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

		responseType := c.Request().FormValue("responseType")

		if isHx.(bool) {
			c.Response().Header().Add("HX-Trigger", "updateUnreadCount, updateArticleList")
			if responseType == "card" {
				return render(c, component.ArticleCard(article))
			} else {
				return render(c, component.ArticleCondensed(article))
			}
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

func (h DeskHandler) DeskPostSetView(c echo.Context) error {
	view := c.Request().FormValue("view")
	ref := c.Request().Referer()
	isHx := c.Get("isHx")

	if view != "" {
		h.session.SetView(c, view)
		c.Set("view", view)
	} else {
		h.session.SetView(c, "card")
		c.Set("view", "card")
	}

	if isHx != nil {
		if isHx.(bool) {
			c.Response().Header().Add("Hx-Redirect", ref)
		}
	}

	return c.Redirect(http.StatusSeeOther, ref)
}

func (h DeskHandler) PostDeskSetReadView(c echo.Context) error {
	readView := c.Request().FormValue("viewRead")
	ref := c.Request().Referer()
	isHx := c.Get("isHx")

	if readView == "true" {
		h.session.SetViewRead(c, true)
	} else {
		h.session.SetViewRead(c, false)
	}

	if isHx != nil {
		if isHx.(bool) {
			c.Response().Header().Add("Hx-Redirect", ref)
		}
	}

	return c.Redirect(http.StatusSeeOther, ref)
}
