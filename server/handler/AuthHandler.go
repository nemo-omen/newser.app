package handler

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo/v4"
	"newser.app/server/service"
	"newser.app/view/pages/auth"
)

// MinChars() returns true if a value contains at least n characters.
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// Matches() returns true if a value matches a provided compiled regular
// expression pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func isEmailValid(e string) bool {
	_, err := mail.ParseAddress(e)
	return err == nil
}

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: authService,
	}
}

func (h AuthHandler) GetLogin(c echo.Context) error {
	return render(c, auth.Login())
}

func (h AuthHandler) GetSignup(c echo.Context) error {
	return render(c, auth.Signup())
}

func (h AuthHandler) PostLogin(c echo.Context) error {
	email := c.Request().FormValue("email")
	pass := c.Request().FormValue("password")
	u, err := h.AuthService.Login(email, pass)
	if err != nil {
		// flash specific errors depending on err
		// type
		SetFlash(c, "error", err.Error())
		return render(c, auth.Login())
	}
	SetFlash(c, "successFlash", fmt.Sprintf("%v logged in successfully!", u.Email))
	SetAuthedSession(c)
	return c.Redirect(http.StatusSeeOther, "/desk/")
}

func (h AuthHandler) PostSignup(c echo.Context) error {

	email := c.Request().FormValue("email")
	pass := c.Request().FormValue("password")
	conf := c.Request().FormValue("confirm")

	if !MinChars(pass, 6) {
		SetFlash(c, "passwordError", "password must be at least 8 characters long")
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	if pass != conf {
		SetFlash(c, "confirmError", "passwords do not match")
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	if !isEmailValid(email) {
		SetFlash(c, "emailError", "not a valid email")
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	pHash := HashPassword(pass)

	u, err := h.AuthService.Signup(email, pHash)
	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "UNIQUE") {
			errmsg = "An account with that email already exisis."
		} else {
			errmsg = err.Error()
		}
		SetFlash(c, "errorFlash", errmsg)
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	SetFlash(c, "successFlash", fmt.Sprintf("%v signed up! Log in to get started.", u.Email))

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (h AuthHandler) PostLogout(c echo.Context) error {
	RevokeAuthedSession(c)
	SetFlash(c, "successFlash", "You are logged out.")
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
