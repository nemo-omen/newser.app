package handler

import (
	"github.com/labstack/echo/v4"
	"newser.app/internal/dto"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/session"
	"newser.app/view/pages/app"
)

type WebAppHandler struct {
	session     session.SessionService
	authService auth.AuthService
	// collectionService collection.CollectionService
	// newsfeedService newsfeed.NewsfeedService
	// subscriptionService subscription.SubscriptionService
	// discoveryService discovery.DiscoveryService
}

func NewWebAppHandler(sessionService session.SessionService, authService auth.AuthService) *WebAppHandler {
	return &WebAppHandler{
		session:     sessionService,
		authService: authService,
		// collectionService: collectionService,
		// newsfeedService: newsfeedService,
		// subscriptionService: subscriptionService,
		// discoveryService: discoveryService,
	}
}

func (h *WebAppHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	app.GET("/app", h.GetApp)
}

func (h *WebAppHandler) GetApp(c echo.Context) error {
	return render(c, app.Index([]*dto.ArticleDTO{}))
}
