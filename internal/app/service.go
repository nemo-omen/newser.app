package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) startService() error {
	// init repos
	// init services
	// init handlers
	app.echo.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong ⚾")
	})
	return nil
}
