package handler

import (
	"current/infra/repository"
	"current/service"
	"current/view/auth"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (h *AuthHandler) HandleGetLogin(c echo.Context) error {
	user := c.Get("user")
	if user != nil {
		return c.Redirect(http.StatusSeeOther, "/app/")
	}
	return render(c, auth.Login())
}

func (h AuthHandler) HandlePostLogin(c echo.Context) error {
	user := c.Get("user")
	if user != nil {
		return c.Redirect(http.StatusSeeOther, "/app/")
	}
	return c.Redirect(http.StatusSeeOther, "/app/")
	// return render(c, auth.Login())
}

func (h *AuthHandler) HandleGetSignup(c echo.Context) error {
	user := c.Get("user")
	if user != nil {
		return c.Redirect(http.StatusSeeOther, "/app/")
	}
	return render(c, auth.Signup())
}

func (h AuthHandler) HandlePostSignup(c echo.Context) error {
	userRepo := repository.NewUserMemRepo()
	userService := service.NewUserService(userRepo)

	email := c.Request().FormValue("email")
	password := c.Request().FormValue("password")
	confirm := c.Request().FormValue("password-confirm")

	if password != confirm {
		// flash error & return signup form
		fmt.Println("passwords do not match")
	}

	u, err := userService.SignUpUser(email, password)
	if err != nil {
		// return signup form w/errors
		fmt.Println("error with user signup: ", err)
	}

	c.Set("user", &u)
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
