package handler

import (
	"current/service"
	"current/view/search"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
)

type SearchHandler struct{}

func (h SearchHandler) HandleGetIndex(c echo.Context) error {
	return render(
		c,
		search.Index(
			search.SearchPageProps{
				Error: "",
				Feeds: []*gofeed.Feed{},
			},
		),
	)
}

func (h SearchHandler) HandlePostSearch(c echo.Context) error {
	searchLinks := []string{}
	api := service.NewAPI(http.DefaultClient)
	formData := c.FormValue("searchurl")
	feeds := []*gofeed.Feed{}

	// Compile a list of feed links
	// 1. Search for links in doc head
	documentFeedLinks, err := api.FindFeedLinks(formData)
	if err != nil {
		// TODO: We should recover and move on to guesses
		// TODO: Session flash
		// TODO: HTMX partial w/form & messages response
		// return render(c, search.Index(search.SearchPageProps{Error: ""}))
		fmt.Println(err)
	}

	searchLinks = append(searchLinks, documentFeedLinks...)

	if len(searchLinks) < 1 {
		guessedLinks, err := api.GuessFeedLinks(formData)
		if err != nil {
			return render(
				c,
				search.Index(
					search.SearchPageProps{
						Error: "Could not find links",
						Feeds: feeds,
					},
				),
			)
		}
		searchLinks = append(searchLinks, guessedLinks...)
	}

	feedsResult, err := api.GetFeedsConcurrent(searchLinks)
	if err != nil {
		return render(
			c,
			search.Index(search.SearchPageProps{
				Error: "could not find any feeds",
				Feeds: feeds,
			}),
		)
	}
	feeds = append(feeds, feedsResult...)

	return render(
		c,
		search.Index(search.SearchPageProps{
			Error: "",
			Feeds: feeds,
		}),
	)
}
