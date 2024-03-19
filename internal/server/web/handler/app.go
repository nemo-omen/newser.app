package handler

import (
	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/session"
	"newser.app/internal/usecase/subscription"
	"newser.app/view/pages/app"
)

type WebAppHandler struct {
	session     session.SessionService
	authService auth.AuthService
	// collectionService collection.CollectionService
	// newsfeedService newsfeed.NewsfeedService
	subscriptionService subscription.SubscriptionService
	// discoveryService discovery.DiscoveryService
}

func NewWebAppHandler(
	sessionService session.SessionService,
	authService auth.AuthService,
	// collectionService collection.CollectionService,
	// newsfeedService newsfeed.NewsfeedService,
	subscriptionService subscription.SubscriptionService,
	// discoveryService discovery.DiscoveryService,
) *WebAppHandler {
	return &WebAppHandler{
		session:     sessionService,
		authService: authService,
		// collectionService: collectionService,
		// newsfeedService: newsfeedService,
		subscriptionService: subscriptionService,
		// discoveryService: discoveryService,
	}
}

func (h *WebAppHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/app", h.GetApp)
	app.GET("/app/article/:id", h.GetArticle)
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

	if isHxRequest(c) {
		return render(c, app.IndexPageContent(articles))
	}

	return render(c, app.Index(articles))
}

func (h *WebAppHandler) GetArticle(c echo.Context) error {
	// get article by id
	// render article page
	return nil
}
