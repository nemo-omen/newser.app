package middleware

import (
	"fmt"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"newser.app/server/service"
)

func SidebarLinks(
	sess *scs.SessionManager,
	subService *service.SubscriptionService,
	authService *service.AuthService,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userEmail := sess.Get(c.Request().Context(), "user")
			if userEmail == nil {
				return next(c)
			}
			user, err := authService.GetUserByEmail(userEmail.(string))
			if err != nil {
				fmt.Println("error getting user for middleware: ", err.Error())
				return next(c)
			}

			feedLinks, err := subService.GetSubscribedFeedsWithLinks(user.Id)
			if err != nil {
				fmt.Println("error getting user for middleware: ", err.Error())
				return next(c)
			}

			c.Set("feedlinks", feedLinks)

			return next(c)
		}
	}
}
