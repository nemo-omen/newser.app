package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"newser.app/internal/dto"
	"newser.app/internal/usecase/collection"
	"newser.app/internal/usecase/session"
	"newser.app/shared/util"
	"newser.app/view/pages/app"
)

type WebCollectionHandler struct {
	session           session.SessionService
	collectionService collection.CollectionService
}

func NewWebCollectionHandler(
	sessionService session.SessionService,
	collectionService collection.CollectionService,
) *WebCollectionHandler {
	return &WebCollectionHandler{
		session:           sessionService,
		collectionService: collectionService,
	}
}

func (h *WebCollectionHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/app/collection/unread", h.GetUnread)
	app.GET("/app/collection/saved", h.GetSaved)

	app.GET("/app/collection/:id", h.GetCollection)

	// app.POST("/app/collection", h.PostCollection)
	// app.POST("/app/collection/delete", h.PostDeleteCollection)
}

func (h *WebCollectionHandler) GetCollection(c echo.Context) error {
	collectionID := c.Param("id")
	fmt.Println("collectionID:", collectionID)
	return nil
}

func (h *WebCollectionHandler) GetUnread(c echo.Context) error {
	util.SetPageTitle(c, h.session, "Unread Articles")

	if isHxRequest(c) {
		return render(c, app.IndexPageContent([]*dto.ArticleDTO{}))
	}
	return render(c, app.Index([]*dto.ArticleDTO{}))
}

func (h *WebCollectionHandler) GetSaved(c echo.Context) error {
	util.SetPageTitle(c, h.session, "Saved Articles")
	if isHxRequest(c) {
		return render(c, app.IndexPageContent([]*dto.ArticleDTO{}))
	}
	return render(c, app.Index([]*dto.ArticleDTO{}))
}
