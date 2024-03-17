package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/session"
	"newser.app/shared"
	authview "newser.app/view/pages/auth"
)

type WebAuthHandler struct {
	authService auth.AuthService
	session     session.SessionService
}

func NewAuthWebHandler(authService auth.AuthService, sessionService session.SessionService) *WebAuthHandler {
	return &WebAuthHandler{
		authService: authService,
		session:     sessionService,
	}
}

func (h *WebAuthHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	authGroup := app.Group("/auth")
	for _, m := range middleware {
		authGroup.Use(m)
	}

	authGroup.GET("/register", h.GetRegister)
	authGroup.POST("/register", h.PostRegister)
	authGroup.GET("/login", h.GetLogin)
	authGroup.POST("/login", h.PostLogin)
	authGroup.POST("/logout", h.Logout)
	authGroup.GET("/user", h.User)
}

func renderOrRedirect(c echo.Context, component templ.Component, redirectURL string) error {
	if isHxRequest(c) {
		return render(c, component)
	}
	return c.Redirect(http.StatusSeeOther, redirectURL)
}

func (h *WebAuthHandler) GetRegister(c echo.Context) error {
	fmt.Println("GET /auth/register")
	return render(c, authview.Register())
}

func (h *WebAuthHandler) PostRegister(c echo.Context) error {
	fmt.Println("POST /auth/register")
	email := c.Request().FormValue("email")
	name := c.Request().FormValue("name")
	password := c.Request().FormValue("password")
	confirm := c.Request().FormValue("confirm")
	if password != confirm {
		h.session.SetFlash(c, "passwordError", "Passwords do not match")
		return renderOrRedirect(c, authview.RegisterPageContent(), "/auth/register")
	}
	_, err := h.authService.Register(email, name, password)
	if err != nil {
		appErr, ok := err.(*shared.AppError)
		if ok {
			if appErr.ErrType == "value.Email" {
				fmt.Println("invalid email")
				h.session.SetFlash(c, "emailError", appErr.Msg)
			}
			if appErr.ErrType == "value.Name" {
				fmt.Println("invalid name")
				h.session.SetFlash(c, "nameError", appErr.Msg)
			}
			if appErr.ErrType == "value.Password" {
				fmt.Println("invalid password")
				h.session.SetFlash(c, "passwordError", appErr.Msg)
			}
			return renderOrRedirect(c, authview.RegisterPageContent(), "/auth/register")
		}
		// TODO: Custom errors for the service/repository layer
		// ErrIncorrectPassword
	}
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (h *WebAuthHandler) GetLogin(c echo.Context) error {
	fmt.Println("GET /auth/login")
	return render(c, authview.Login())
}

func (h *WebAuthHandler) PostLogin(c echo.Context) error {
	fmt.Println("POST /auth/login")
	email := c.Request().FormValue("email")
	password := c.Request().FormValue("password")
	user, err := h.authService.Login(email, password)
	if err != nil {
		appErr, ok := err.(*shared.AppError)
		if ok {
			if appErr.ErrType == "value.Email" {
				fmt.Println("invalid email")
				h.session.SetFlash(c, "emailError", appErr.Msg)
			}
			if appErr.ErrType == "value.Password" {
				fmt.Println("invalid password")
				h.session.SetFlash(c, "passwordError", appErr.Msg)
			}
			return renderOrRedirect(c, authview.LoginPageContent(), "/auth/login")
		}
	}
	return c.JSON(http.StatusOK, user)
}

func (h *WebAuthHandler) Logout(c echo.Context) error {
	// redirect => "/"
	h.session.RevokeAuth(c)
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (h *WebAuthHandler) User(c echo.Context) error {
	return c.String(http.StatusOK, "user")
}
