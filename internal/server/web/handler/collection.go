package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"newser.app/internal/dto"
	"newser.app/internal/usecase/auth"
	"newser.app/internal/usecase/collection"
	"newser.app/internal/usecase/session"
	"newser.app/internal/usecase/subscription"
	"newser.app/shared/util"
	"newser.app/view/component"
	"newser.app/view/pages/app"
	collectionview "newser.app/view/pages/app/collection"
)

type WebCollectionHandler struct {
	session             session.SessionService
	collectionService   collection.CollectionService
	authService         auth.AuthService
	subscriptionService subscription.SubscriptionService
}

func NewWebCollectionHandler(
	sessionService session.SessionService,
	collectionService collection.CollectionService,
	authService auth.AuthService,
	subscriptionService subscription.SubscriptionService,
) *WebCollectionHandler {
	return &WebCollectionHandler{
		session:             sessionService,
		collectionService:   collectionService,
		authService:         authService,
		subscriptionService: subscriptionService,
	}
}

func (h *WebCollectionHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/app/collections", h.GetCollections)
	app.GET("/app/collection/unread", h.GetUnread)
	app.GET("/app/collection/saved", h.GetSaved)
	app.POST("/app/collection/unread", h.PostUnread)
	app.POST("/app/collection/read", h.PostRead)
	app.POST("/app/collection/saved", h.PostSaved)
	app.POST("/app/collection/unsaved", h.PostUnsaved)
	app.GET("/app/collection/:id", h.GetCollection)
	app.GET("/app/collection/new", h.GetCreateCollection)
	app.POST("/app/collection/new", h.PostCreateCollection)

	// app.POST("/app/collection", h.PostCollection)
	// app.POST("/app/collection/delete", h.PostDeleteCollection)
}

func (h *WebCollectionHandler) GetCollections(c echo.Context) error {
	util.SetPageTitle(c, h.session, "Your Collections")
	if isHxRequest(c) {
		return render(c, collectionview.IndexPageContent([]*dto.CollectionDTO{}))
	}
	return render(c, collectionview.Index([]*dto.CollectionDTO{}))
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
		return render(c, app.IndexPageContent(collectionArticles, "/app/collection/unread"))
	}
	return render(c, app.Index(collectionArticles, "/app/collection/unread"))
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
	// fmt.Println("saved collectionArticles:", collectionArticles)

	util.SetPageTitle(c, h.session, "Saved Articles")
	if isHxRequest(c) {
		return render(c, app.IndexPageContent(collectionArticles, "/app/collection/saved"))
	}
	return render(c, app.Index(collectionArticles, "/app/collection/saved"))
}

func (h *WebCollectionHandler) GetCreateCollection(c echo.Context) error {
	isAuthed, ok := c.Get("authenticated").(bool)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	if !isAuthed {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	util.SetPageTitle(c, h.session, "Create a Collection")
	if isHxRequest(c) {
		return render(c, collectionview.NewPageContent())
	}
	return render(c, collectionview.New())
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
	viewType := c.FormValue("viewtype")
	ref := c.Request().Referer()
	if articleID == "" {
		h.session.SetFlash(c, "error", "Article not found")
		// return redirectWithHX(c, ref)
	}
	err = h.collectionService.AddAndRemoveArticleFromCollection("unread", "read", articleID, user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "Error marking article as unread")
	}
	article, err := h.subscriptionService.GetArticle(user.ID.String(), articleID)
	if err != nil {
		h.session.SetFlash(c, "error", "Error fetching article")
	}

	if isHxRequest(c) {
		c.Response().Header().Set("Hx-Trigger", "update-articles")
		if viewType == "condensed" {
			return render(c, component.ArticleCondensed(article))
		} else if viewType == "expanded" {
			return render(c, component.ArticleCard(article))
		}
	}
	return c.Redirect(http.StatusSeeOther, ref)
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
	viewType := c.FormValue("viewtype")
	ref := c.Request().Referer()
	if articleID == "" {
		h.session.SetFlash(c, "error", "Article not found")
		// return redirectWithHX(c, ref)
	}
	err = h.collectionService.AddAndRemoveArticleFromCollection("read", "unread", articleID, user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "Error marking article as read")
	}
	article, err := h.subscriptionService.GetArticle(user.ID.String(), articleID)
	if err != nil {
		h.session.SetFlash(c, "error", "Error fetching article")
	}

	if isHxRequest(c) {
		c.Response().Header().Set("Hx-Trigger", "update-articles")
		if viewType == "condensed" {
			return render(c, component.ArticleCondensed(article))
		} else if viewType == "expanded" {
			return render(c, component.ArticleCard(article))
		}
	}
	return c.Redirect(http.StatusSeeOther, ref)
}

func (h *WebCollectionHandler) PostSaved(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	articleID := c.FormValue("articleid")
	viewType := c.FormValue("viewtype")
	ref := c.Request().Referer()
	if articleID == "" {
		h.session.SetFlash(c, "error", "Article not found")
		// return redirectWithHX(c, ref)
	}
	err = h.collectionService.AddArticleToCollection("saved", articleID, user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "Error saving article")
	}
	article, err := h.subscriptionService.GetArticle(user.ID.String(), articleID)
	if err != nil {
		h.session.SetFlash(c, "error", "Error fetching article")
	}
	if isHxRequest(c) {
		c.Response().Header().Set("Hx-Trigger", "update-articles")
		if viewType == "condensed" {
			return render(c, component.ArticleCondensed(article))
		} else if viewType == "expanded" {
			return render(c, component.ArticleCard(article))
		}
	}
	return c.Redirect(http.StatusSeeOther, ref)
}

func (h *WebCollectionHandler) PostUnsaved(c echo.Context) error {
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/login")
	}
	articleID := c.FormValue("articleid")
	viewType := c.FormValue("viewtype")
	ref := c.Request().Referer()
	if articleID == "" {
		h.session.SetFlash(c, "error", "Article not found")
		// return redirectWithHX(c, ref)
	}
	err = h.collectionService.RemoveArticleFromCollection("saved", articleID, user.ID.String())
	// err = h.collectionService.AddArticleToCollection("saved", articleID, user.ID.String())
	if err != nil {
		h.session.SetFlash(c, "error", "Error saving article")
	}
	article, err := h.subscriptionService.GetArticle(user.ID.String(), articleID)
	if err != nil {
		h.session.SetFlash(c, "error", "Error fetching article")
	}
	if isHxRequest(c) {
		c.Response().Header().Set("Hx-Trigger", "update-articles")
		if viewType == "condensed" {
			return render(c, component.ArticleCondensed(article))
		} else if viewType == "expanded" {
			return render(c, component.ArticleCard(article))
		}
	}
	return c.Redirect(http.StatusSeeOther, ref)
}

func (h *WebCollectionHandler) PostCreateCollection(c echo.Context) error {
	collectionName := c.FormValue("name")
	collectionDescription := c.FormValue("description")
	email, ok := c.Get("user").(string)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	user, err := h.authService.GetUserByEmail(email)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	// err := h.collectionService.
	fmt.Println(user)
	return c.Redirect(http.StatusSeeOther, "/")
}
