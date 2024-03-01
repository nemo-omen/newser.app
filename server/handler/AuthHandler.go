package handler

import (
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/alexedwards/scs/v2"
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
	session     Session
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService, sessionManager *scs.SessionManager) AuthHandler {
	return AuthHandler{
		session:     Session{manager: sessionManager},
		authService: authService,
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
	_, err := h.authService.Login(email, pass)
	if err != nil {
		// flash specific errors depending on err
		// type
		h.session.SetFlash(c, "error", err.Error())
		return render(c, auth.Login())
	}
	h.session.SetFlash(c, "successFlash", fmt.Sprintf("%v logged in successfully!", email))
	h.session.SetAuth(c, email)
	return c.Redirect(http.StatusSeeOther, "/desk/")
}

func (h AuthHandler) PostSignup(c echo.Context) error {

	email := c.Request().FormValue("email")
	pass := c.Request().FormValue("password")
	conf := c.Request().FormValue("confirm")

	if !MinChars(pass, 6) {
		h.session.SetFlash(c, "passwordError", "password must be at least 8 characters long")
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	if pass != conf {
		h.session.SetFlash(c, "confirmError", "passwords do not match")
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	if !isEmailValid(email) {
		h.session.SetFlash(c, "emailError", "not a valid email")
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	pHash := HashPassword(pass)

	u, err := h.authService.Signup(email, pHash)
	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "UNIQUE") {
			errmsg = "An account with that email already exisis."
		} else {
			errmsg = err.Error()
		}
		h.session.SetFlash(c, "errorFlash", errmsg)
		return c.Redirect(http.StatusSeeOther, "/auth/signup")
	}

	h.session.SetFlash(c, "successFlash", fmt.Sprintf("%v signed up! Log in to get started.", u.Email))

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (h AuthHandler) PostLogout(c echo.Context) error {
	h.session.RevokeAuth(c)
	h.session.SetFlash(c, "successFlash", "You are logged out.")
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
