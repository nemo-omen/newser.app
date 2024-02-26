package handler

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"unicode/utf8"

	"github.com/labstack/echo/v4"
	"newser.app/internal/service"
	"newser.app/ui/view/pages/auth"
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

func NewAuthHandler(dsn string) AuthHandler {
	return AuthHandler{
		AuthService: service.NewAuthService(dsn),
	}
}

func (h AuthHandler) GetLogin(c echo.Context) error {
	_ = GetFlash(c, "errorFlash")
	_ = GetFlash(c, "successFlash")
	_ = GetFlash(c, "notificationFlash")

	return render(c, auth.Login())
}

func (h AuthHandler) GetSignup(c echo.Context) error {
	_ = GetFlash(c, "emailError")
	_ = GetFlash(c, "passwordError")
	_ = GetFlash(c, "confirmError")

	return render(c, auth.Signup())
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
		SetFlash(c, "errorFlash", err.Error())
	}

	SetFlash(c, "successFlash", fmt.Sprintf("%v signed up! Log in to get started.", u.Email))

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
