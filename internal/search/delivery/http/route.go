package http

import (
	"github.com/labstack/echo/v4"
)

func (h *SearchHandler) Routes(group *echo.Group) {
	// route: POST /api/v1/search
	group.POST("/search", h.PostSearch)

	// middleware example
	// group.GET("/whatever", h.GetWhatever, middleware.WhateverMiddleware())
}
