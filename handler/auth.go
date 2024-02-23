package handler

import (
	"current/infra/repository"
	"current/service"
	"current/view/auth"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
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
	flash := &auth.SignupFlash{}
	user := c.Get("user")
	if user != nil {
		return c.Redirect(http.StatusSeeOther, "/app/")
	}
	return render(c, auth.Signup(flash))
}

func (h AuthHandler) HandlePostSignup(c echo.Context) error {
	flash := &auth.SignupFlash{}
	userRepo := repository.NewUserMemRepo()
	userService := service.NewUserService(userRepo)
	email := c.Request().FormValue("email")
	password := c.Request().FormValue("password")
	confirm := c.Request().FormValue("password-confirm")
	sess, _ := session.Get("session", c)
	// sess.Options = &sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   7 * int(time.Duration.Hours(24)),
	// 	HttpOnly: true,
	// 	SameSite: http.SameSiteStrictMode,
	// }

	fmt.Println(sess)

	if password != confirm {
		// flash error & return signup form
		fmt.Println("passwords do not match")
		flash.ConfirmError = "Passwords do not match"
		return render(c, auth.Signup(flash))
	}

	if len(password) < 8 {
		flash.PasswordError = "Password must be at least 8 characters long"
		return render(c, auth.Signup(flash))
	}

	u, err := userService.SignUpUser(email, password)
	if err != nil {
		// return signup form w/errors
		fmt.Println("error with user signup: ", err)
		flash.GlobalError = err.Error()
		return render(c, auth.Signup(flash))
	}
	sess.Values["k"] = "fakeKeyOhYeah"

	sess.Save(c.Request(), c.Response())
	fmt.Println(u)
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
