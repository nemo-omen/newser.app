package service

import (
	"current/util"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"current/common"

	"github.com/PuerkitoBio/goquery"
)

type API struct {
	Client *http.Client
}

func NewAPI(c *http.Client) *API {
	return &API{
		Client: c,
	}
}

func (api *API) FindFeedLinks(siteUrl string) ([]string, error) {
	links := []string{}
	safeUrl, err := util.MakeUrl(siteUrl)
	if err != nil {
		return links, err
	}

	res, err := api.Client.Get(safeUrl.String())
	if err != nil {
		return links, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return links, fmt.Errorf("request returned error status: %+v", res.Status)
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return links, err
	}

	document.Find("link").Each(func(i int, s *goquery.Selection) {
		rel, exists := s.Attr("rel")
		if !exists {
			return
		}

		if rel != "alternate" {
			return
		}

		linkType, exists := s.Attr("type")

		if !exists {
			return
		}

		if !slices.Contains(common.ValidFeedLinkTypes, linkType) {
			return
		}

		href, exists := s.Attr("href")
		if !exists {
			return
		}

		testUrl, err := url.Parse(href)
		if err != nil {
			return
		}

		if !testUrl.IsAbs() {
			testUrl.Scheme = safeUrl.Scheme
			testUrl.Host = safeUrl.Host
			links = append(links, testUrl.String())
		} else {
			links = append(links, href)
		}
	})

	return links, nil
}
