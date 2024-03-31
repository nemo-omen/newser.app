package handler

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
	"newser.app/internal/dto"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/collection"
	"newser.app/internal/usecase/discovery"
	"newser.app/internal/usecase/newsfeed"
	"newser.app/internal/usecase/session"
	"newser.app/internal/usecase/subscription"
	"newser.app/shared/util"
	"newser.app/view/component"
	"newser.app/view/pages/app"
)

type WebAppHandler struct {
	session             session.SessionService
	authService         auth.AuthService
	subscriptionService subscription.SubscriptionService
	collectionService   collection.CollectionService
	feedApi             discovery.DiscoveryService
	newsfeedService     newsfeed.NewsfeedService
	// discoveryService discovery.DiscoveryService
}

func NewWebAppHandler(
	sessionService session.SessionService,
	authService auth.AuthService,
	subscriptionService subscription.SubscriptionService,
	collectionService collection.CollectionService,
	feedApi discovery.DiscoveryService,
	newsfeedService newsfeed.NewsfeedService,
	// discoveryService discovery.DiscoveryService,
) *WebAppHandler {
	return &WebAppHandler{
		session:             sessionService,
		authService:         authService,
		subscriptionService: subscriptionService,
		collectionService:   collectionService,
		feedApi:             feedApi,
		newsfeedService:     newsfeedService,
		// discoveryService: discoveryService,
	}
}

func (h *WebAppHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/app", h.GetApp)
	app.GET("/app/newsfeed/:id", h.GetNewsfeed)
	app.GET("/app/article/:id", h.GetArticle)
	app.GET("/app/control/unreadcount", h.GetUpdatedSidebarCount)
	app.GET("/app/control/currentpath", h.GetUpdatedSidebar)
	app.GET("/app/control/pagetitle", h.PageTitle)
	app.POST("/app/control/viewread", h.ViewRead)
	app.POST("/app/control/viewunread", h.ViewUnread)
	app.POST("/app/control/viewcollapsed", h.ViewCondensed)
	app.POST("/app/control/viewexpanded", h.ViewExpanded)
}

func (h *WebAppHandler) GetApp(c echo.Context) error {
	c.Set("title", "Latest Articles")
	authed, ok := c.Get("authenticated").(bool)
	if !ok {
		authed = false
	}
	if !authed {
		return redirectWithHX(c, "/auth/login")
	}
	// get all the subscribed articles
	email, ok := c.Get("user").(string)
	if !ok {
		return redirectWithHX(c, "/auth/login")
	}

	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return redirectWithHX(c, "/auth/login")
	}

	articles, err := h.subscriptionService.GetAllArticles(user.ID.String())
	if err != nil {
		// set flash message
		// render or redirect to error page? /app?
	}
	// articles := []*dto.ArticleDTO{}
	if len(articles) == 0 {
		// renderOrRedirect(c, search.SearchPageContent([]*gofeed.Feed{}), "/app/search")
		return redirectWithHX(c, "/app/search")
	}

	util.SetPageTitle(c, h.session, "Latest Articles")

	if isHxRequest(c) {
		return render(c, app.IndexPageContent(articles, "/app"))
	}
	return render(c, app.Index(articles, "/app"))
}

func (h *WebAppHandler) UpdateFeeds(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	currentArticles, err := h.subscriptionService.GetAllArticles(user.ID.String())

	if err != nil {
		h.session.SetFlash(c, "error", "There was a problem getting stored articles")
	}

	feeds, err := h.subscriptionService.GetAllFeeds(user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "There was a problem updating your feeds.")
	}
	feedUrls := []string{}
	articleUrls := []string{}

	for _, feed := range feeds {
		feedUrls = append(feedUrls, feed.FeedURL)
		for _, article := range feed.Articles {
			articleUrls = append(articleUrls, article.Link)
		}
	}
	fmt.Println("currentArticles: ", currentArticles)

	goFeeds, err := h.feedApi.GetFeedsConcurrent(feedUrls)

	for _, goFeed := range goFeeds {
		for _, item := range goFeed.Items {
			if !slices.Contains(articleUrls, item.Link) {
				err = h.newsfeedService.SaveArticle(item)
				if err != nil {
					h.session.SetFlash(c, "error", "There was a problem saving some articles.")
				}
			}
		}
	}

	fmt.Println(goFeeds)
	// TODO: we're going to need another, smaller component
	// to send back that just has updated, sorted articles
	return render(c, app.IndexPageContent([]*dto.ArticleDTO{}, "/app"))
}

func (h *WebAppHandler) GetNewsfeed(c echo.Context) error {
	// get newsfeed by id
	// render newsfeed page
	feedId := c.Param("id")
	email, ok := c.Get("user").(string)
	if !ok {
		return redirectWithHX(c, "/auth/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return redirectWithHX(c, "/auth/login")
	}
	feed, err := h.subscriptionService.GetNewsfeed(user.ID.String(), feedId)
	if err != nil {
		h.session.SetFlash(c, "error", "Failed to get newsfeed")
		return redirectWithHX(c, "/app")
		// set flash message
		// render or redirect to error page? /app?
	}
	util.SetPageTitle(c, h.session, feed.Title)

	if isHxRequest(c) {
		return render(c, app.FeedPageContent(feed))
	}

	return render(c, app.Feed(feed))
}

func (h *WebAppHandler) GetArticle(c echo.Context) error {
	articleId := c.Param("id")
	email, ok := c.Get("user").(string)
	if !ok {
		return redirectWithHX(c, "/auth/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return redirectWithHX(c, "/auth/login")
	}
	article, err := h.subscriptionService.GetArticle(user.ID.String(), articleId)
	if err != nil {
		h.session.SetFlash(c, "error", "Failed to get newsfeed")
		return redirectWithHX(c, "/app")
		// set flash message
		// render or redirect to error page? /app?
	}

	// if !article.Read {
	// 	article.Read = true
	// 	_ = h.collectionService.AddAndRemoveArticleFromCollection(
	// 		"read",
	// 		"unread",
	// 		article.ID.String(),
	// 		user.ID.String(),
	// 	)
	// }

	util.SetPageTitle(c, h.session, article.Title)

	if isHxRequest(c) {
		return render(c, app.ArticlePageContent(article))
	}

	return render(c, app.Article(article))
}

func (h *WebAppHandler) GetUpdatedSidebarCount(c echo.Context) error {
	return render(c, component.MainFeedLinks())
}

func (h *WebAppHandler) GetUpdatedSidebar(c echo.Context) error {
	return render(c, component.MainSidebar())
}

func (h *WebAppHandler) PageTitle(c echo.Context) error {
	title := c.Get("title")
	titleString, ok := title.(string)
	// fmt.Println("titleString: ", titleString)
	if !ok {
		return c.String(http.StatusOK, "Newser")
	}
	return c.String(http.StatusOK, titleString)
}

func (h *WebAppHandler) ViewUnread(c echo.Context) error {
	ref := c.Request().Referer()
	h.session.SetView(c, "unread")
	if isHxRequest(c) {
		c.Response().Header().Set("HX-Redirect", ref)
	}
	return c.Redirect(http.StatusSeeOther, ref)
}

func (h *WebAppHandler) ViewRead(c echo.Context) error {
	ref := c.Request().Referer()
	h.session.SetView(c, "read")
	if isHxRequest(c) {
		c.Response().Header().Set("HX-Redirect", ref)
	}
	return c.Redirect(http.StatusSeeOther, ref)
}

func (h *WebAppHandler) ViewCondensed(c echo.Context) error {
	ref := c.Request().Referer()
	h.session.SetLayout(c, "condensed")
	if isHxRequest(c) {
		c.Response().Header().Set("HX-Redirect", ref)
	}
	return c.Redirect(http.StatusSeeOther, ref)
}

func (h *WebAppHandler) ViewExpanded(c echo.Context) error {
	ref := c.Request().Referer()
	h.session.SetLayout(c, "expanded")
	if isHxRequest(c) {
		c.Response().Header().Set("HX-Redirect", ref)
	}
	return c.Redirect(http.StatusSeeOther, ref)
}
