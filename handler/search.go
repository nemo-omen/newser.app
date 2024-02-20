package handler

import (
	"current/service"
	"current/util"
	"current/view"
	"current/view/component"
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
			view.SearchPageProps{
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
	isHX := c.Get("isHX").(bool)
	// Make sure we have a valid url
	validUrl, err := util.MakeUrl(formData)
	searchUrl := validUrl.String()

	if err != nil {
		return render(
			c,
			search.Index(
				view.SearchPageProps{
					Error: fmt.Sprintf("Not a valid url, try %v.com?", formData),
					Feeds: feeds,
				},
			),
		)
	}

	// make sure a site exists at validUrl,
	isSite := api.CheckSite(searchUrl)
	if !isSite {
		return render(
			c,
			search.Index(
				view.SearchPageProps{
					Error: fmt.Sprintf("Could not find a site at %v", searchUrl),
					Feeds: feeds,
				},
			),
		)
	}

	// we should safely be able to start checking
	// for links

	// Compile a list of feed links
	// 1. Search for links in doc head
	documentFeedLinks, err := api.FindFeedLinks(searchUrl)
	if err != nil {
		// this is just the first step in establishing
		// whether we have any good feed links, so I don't
		// want to render anything yet
		fmt.Println(err)
	}

	searchLinks = append(searchLinks, documentFeedLinks...)

	// 2. No doc links? Try guessing.
	if len(searchLinks) < 1 {
		guessedLinks, err := api.GuessFeedLinks(searchUrl)
		if err != nil {
			return render(
				c,
				search.Index(
					view.SearchPageProps{
						Error: "Could not find links",
						Feeds: feeds,
					},
				),
			)
		}
		searchLinks = append(searchLinks, guessedLinks...)
	}

	// 3. Get feeds from all links concurrently
	feedsResult, err := api.GetFeedsConcurrent(searchLinks)
	if err != nil {
		return render(
			c,
			search.Index(view.SearchPageProps{
				Error: fmt.Sprintf("could not find any feeds at %v", formData),
				Feeds: feeds,
			}),
		)
	}
	feeds = append(feeds, feedsResult...)

	if isHX {
		return render(
			c,
			component.SearchResult(feeds),
		)
	}
	return render(
		c,
		search.Index(view.SearchPageProps{
			Error: "",
			Feeds: feeds,
		}),
	)
}
