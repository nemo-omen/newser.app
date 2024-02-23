package handler

import (
	"current/service"
	"current/view/subscription"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SubscriptionHandler struct{}

func (h SubscriptionHandler) HandleGetIndex(c echo.Context) error {
	return render(c, subscription.Index())
}

func (h SubscriptionHandler) HandlePostSubscribe(c echo.Context) error {
	api := service.NewAPI(&http.Client{})
	subscriptionUrl := c.Request().FormValue("subscriptionurl")
	feed, err := api.GetFeed(subscriptionUrl)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/subscriptions/")
	}
	fmt.Println(feed)
	return c.Redirect(http.StatusSeeOther, "/app/subscriptions/")
}
