package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/collection"
	"newser.app/internal/usecase/session"
	"newser.app/shared/util"
	"newser.app/view/pages/app"
)

type WebCollectionHandler struct {
	session           session.SessionService
	collectionService collection.CollectionService
	authService       auth.AuthService
}

func NewWebCollectionHandler(
	sessionService session.SessionService,
	collectionService collection.CollectionService,
	authService auth.AuthService,
) *WebCollectionHandler {
	return &WebCollectionHandler{
		session:           sessionService,
		collectionService: collectionService,
		authService:       authService,
	}
}

func (h *WebCollectionHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/app/collection/unread", h.GetUnread)
	app.GET("/app/collection/saved", h.GetSaved)
	app.POST("/app/collection/unread", h.PostUnread)
	app.POST("/app/collection/read", h.PostRead)
	// app.POST("/app/collection/saved", h.PostSaved)

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
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}

	collectionArticles, err := h.collectionService.GetArticlesBySlug("unread", user.ID.String())
	if err != nil {
		fmt.Println("error fetching unread articles: ", err.Error())
		h.session.SetFlash(c, "error", "Error fetching unread articles")
		return redirectWithHX(c, "/app/collection/unread")
	}

	util.SetPageTitle(c, h.session, "Unread Articles")
	if isHxRequest(c) {
		return render(c, app.IndexPageContent(collectionArticles))
	}
	return render(c, app.Index(collectionArticles))
}

func (h *WebCollectionHandler) GetSaved(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}

	collectionArticles, err := h.collectionService.GetArticlesBySlug("saved", user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "Error fetching saved articles")
		fmt.Println("error fetching saved articles: ", err.Error())
		return redirectWithHX(c, "/app/collection/saved")
	}
	fmt.Println("saved collectionArticles:", collectionArticles)

	util.SetPageTitle(c, h.session, "Saved Articles")
	if isHxRequest(c) {
		return render(c, app.IndexPageContent(collectionArticles))
	}
	return render(c, app.Index(collectionArticles))
}

func (h *WebCollectionHandler) PostUnread(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	articleID := c.FormValue("articleid")
	ref := c.Request().Referer()
	if articleID == "" {
		h.session.SetFlash(c, "error", "Article not found")
		return redirectWithHX(c, ref)
	}
	err = h.collectionService.AddAndRemoveArticleFromCollection("unread", "read", articleID, user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "Error marking article as unread")
	}
	return redirectWithHX(c, ref)
}

func (h *WebCollectionHandler) PostRead(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	articleID := c.FormValue("articleid")
	ref := c.Request().Referer()
	if articleID == "" {
		h.session.SetFlash(c, "error", "Article not found")
		return redirectWithHX(c, ref)
	}
	err = h.collectionService.AddAndRemoveArticleFromCollection("read", "unread", articleID, user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "Error marking article as read")
	}
	return redirectWithHX(c, ref)
}
