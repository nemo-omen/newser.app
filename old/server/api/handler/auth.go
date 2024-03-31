package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/auth"
)

type ApiAuthHandler struct {
	authService auth.AuthService
}

func NewAuthApiHandler(authService auth.AuthService) *ApiAuthHandler {
	return &ApiAuthHandler{
		authService: authService,
	}
}

func (h *ApiAuthHandler) Routes(app *echo.Echo) {
	authGroup := app.Group("/api/auth")
	authGroup.POST("/register", h.PostRegister)
	authGroup.POST("/login", h.PostLogin)
	authGroup.POST("/logout", h.PostLogout)
	authGroup.GET("/user", h.GetUser)
}

func (h *ApiAuthHandler) PostRegister(c echo.Context) error {
	fmt.Println("POST /api/auth/register")
	// want: JSON body with email, name, password
	return c.JSON(http.StatusOK, "register")
}

func (h *ApiAuthHandler) PostLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, "login")
}

func (h *ApiAuthHandler) PostLogout(c echo.Context) error {
	return c.JSON(http.StatusOK, "logout")
}

func (h *ApiAuthHandler) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "user")
}
