package app

import (
	"net/http"

	searchHttp "newser.app/internal/search/delivery/http"
	searchRepo "newser.app/internal/search/repository"
	searchUsecase "newser.app/internal/search/usecase"

	"github.com/labstack/echo/v4"
)

func (app *App) startService() error {
	// init repos
	searchRepo := searchRepo.NewGofeedRepository(&http.Client{})
	// init services
	searchService := searchUsecase.NewSearchService(searchRepo)
	// init handlers
	searchHandler := searchHttp.NewSearchHandler(searchService)

	searchGroup := app.echo.Group("/search")
	app.echo.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong ⚾")
	})
	searchHandler.Routes(searchGroup)
	return nil
}
