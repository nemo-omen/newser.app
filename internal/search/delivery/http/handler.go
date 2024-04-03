package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"newser.app/internal/search"
	"newser.app/internal/search/dto"
)

// SearchHandler provides HTTP handlers for search operations.
type SearchHandler struct {
	service search.SearchUsecase
}

// NewSearchHandler creates a new SearchHandler.
func NewSearchHandler(service search.SearchUsecase) *SearchHandler {
	return &SearchHandler{
		service: service,
	}
}

func (h *SearchHandler) PostSearch(c echo.Context) error {
	var req dto.SearchRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	feedLinks, err := h.service.FindFeedUrls(req.SearchUrl)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	feeds, err := h.service.GetFeeds(feedLinks)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, feeds)
}
