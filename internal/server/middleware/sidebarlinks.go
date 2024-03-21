package middleware

import (
	"fmt"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/subscription"
)

func SidebarLinks(
	sess *scs.SessionManager,
	subService *subscription.SubscriptionService,
	authService *auth.AuthService,
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

			feedLinks, err := subService.GetSidebarLinks(user.ID.String())
			if err != nil {
				fmt.Println("error getting feeds for middleware: ", err.Error())
				return next(c)
			}

			// fmt.Println("feedLinks: ", feedLinks)
			c.Set("feedlinks", feedLinks)

			return next(c)
		}
	}
}
