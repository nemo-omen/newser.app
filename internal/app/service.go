package app

import (
	"net/http"

	searchHttp "newser.app/internal/search/delivery/http"
	searchRepo "newser.app/internal/search/repository"
	searchUsecase "newser.app/internal/search/usecase"

	"github.com/labstack/echo/v4"
)

// repos
var (
	searchRepository *searchRepo.GofeedRepository
)

// usecases
var (
	searchService *searchUsecase.SearchService
)

// handlers
var (
	apiSearchHandler *searchHttp.SearchHandler
)

func (app *App) startService() error {
	initRepositories()
	initServices()
	initHandlers()

	apiGroup := app.echo.Group("/api/v1")
	apiSearchHandler.Routes(apiGroup)

	app.echo.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong ⚾")
	})
	return nil
}

func initRepositories() {
	searchRepository = searchRepo.NewGofeedRepository(&http.Client{})
}

func initServices() {
	searchService = searchUsecase.NewSearchService(searchRepository)
}

func initHandlers() {
	apiSearchHandler = searchHttp.NewSearchHandler(searchService)
}
