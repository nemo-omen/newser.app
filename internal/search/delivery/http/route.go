package http

import (
	"github.com/labstack/echo/v4"
)

func (h *SearchHandler) Routes(group *echo.Group) {
	// route: POST /search
	group.POST("/feed", h.PostSearch)

	// middleware example
	// group.GET("/whatever", h.GetWhatever, middleware.WhateverMiddleware())
}
