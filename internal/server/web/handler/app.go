package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/auth"
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
	// collectionService collection.CollectionService
	// newsfeedService     newsfeed.NewsfeedService
	// discoveryService discovery.DiscoveryService
}

func NewWebAppHandler(
	sessionService session.SessionService,
	authService auth.AuthService,
	subscriptionService subscription.SubscriptionService,
	// collectionService collection.CollectionService,
	// newsfeedService newsfeed.NewsfeedService,
	// discoveryService discovery.DiscoveryService,
) *WebAppHandler {
	return &WebAppHandler{
		session:             sessionService,
		authService:         authService,
		subscriptionService: subscriptionService,
		// collectionService: collectionService,
		// newsfeedService:     newsfeedService,
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
		return render(c, app.IndexPageContent(articles))
	}
	return render(c, app.Index(articles))
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
