package handler

import (
	"github.com/labstack/echo/v4"
	"newser.app/internal/usecase/session"
	"newser.app/view/pages/note"
)

type WebNoteHandler struct {
	session session.SessionService
}

func NewWebNoteHandler(session session.SessionService) *WebNoteHandler {
	return &WebNoteHandler{
		session: session,
	}
}

func (h *WebNoteHandler) Routes(app *echo.Echo, middleware ...echo.MiddlewareFunc) {
	for _, m := range middleware {
		app.Use(m)
	}
	app.GET("/app/note", h.GetNotes)
}

func (h *WebNoteHandler) GetNotes(c echo.Context) error {
	if isHxRequest(c) {
		return render(c, note.IndexPageContent())
	}
	return render(c, note.Index())
}
