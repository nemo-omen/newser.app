package handler

import (
	"current/view/search"
	"fmt"

	"github.com/labstack/echo/v4"
)

type SearchHandler struct{}

func (h SearchHandler) HandleGetIndex(c echo.Context) error {
	return render(c, search.Index())
}

func (h SearchHandler) HandlePostSearch(c echo.Context) error {
	searchUrl := c.FormValue("searchurl")
	fmt.Println(searchUrl)
	return render(c, search.Index())
}
