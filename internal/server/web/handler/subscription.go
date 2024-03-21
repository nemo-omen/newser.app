package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/discovery"
	"newser.app/internal/usecase/session"
	"newser.app/internal/usecase/subscription"
	"newser.app/shared"
	"newser.app/view/pages/app"
)

type WebSubscriptionHandler struct {
	session             session.SessionService
	subscriptionService subscription.SubscriptionService
	authService         auth.AuthService
	searchService       discovery.DiscoveryService
}

func NewWebSubscriptionHandler(
	sessionService session.SessionService,
	subscriptionService subscription.SubscriptionService,
	authService auth.AuthService,
	searchService discovery.DiscoveryService,
) *WebSubscriptionHandler {
	return &WebSubscriptionHandler{
		session:             sessionService,
		subscriptionService: subscriptionService,
		authService:         authService,
		searchService:       searchService,
	}
}

func (h *WebSubscriptionHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.POST("/app/subscribe", h.PostSubscribe)
	app.POST("/app/unsubscribe", h.PostUnSubscribe)
}

func (h *WebSubscriptionHandler) PostSubscribe(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return redirectWithHX(c, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return redirectWithHX(c, "/app/login")
	}
	userID := user.ID.String()

	subscriptionUrl := c.FormValue("subscriptionurl")
	gofeed, err := h.searchService.GetFeed(subscriptionUrl)
	if err != nil {
		appErr, ok := err.(shared.AppError)
		if ok {
			appErr.Print()
		}
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error fetching feed: %v", err))
	}
	subscription, err := h.subscriptionService.Subscribe(userID, *gofeed)

	if err != nil {
		appErr, ok := err.(shared.AppError)
		if ok {
			fmt.Println("err: ", appErr.Err)
			fmt.Println("errType: ", appErr.ErrType)
			fmt.Println("errMsg: ", appErr.Msg)
			fmt.Println("errOrigin: ", appErr.Origin)
		}
		h.session.SetFlash(c, "error", "Error subscribing to feed")
		return c.Redirect(http.StatusSeeOther, "/app/search")
	}

	feed, err := h.subscriptionService.GetNewsfeed(userID, subscription.FeedID)
	if err != nil {
		h.session.SetFlash(c, "error", "Error fetching feed")
		return c.Redirect(http.StatusSeeOther, "/app/search")
	}

	h.session.SetFlash(c, "success", "Subscribed to feed")
	return renderOrRedirect(c, app.FeedPageContent(feed), fmt.Sprintf("/app/newsfeed/%s", feed.ID))
}

func (h *WebSubscriptionHandler) PostUnSubscribe(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}

	feedID := c.FormValue("feedID")
	err = h.subscriptionService.UnSubscribe(user.ID.String(), feedID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Unsubscribed from %v", feedID))
}
