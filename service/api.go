package service

import (
	"current/util"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"current/common"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
)

type API struct {
	Client *http.Client
}

func NewAPI(c *http.Client) *API {
	return &API{
		Client: c,
	}
}

// GetFeed attempts to retrieve a valid RSS/Atom/JSON feed
// from the given URL string.
//
// The URL string must include a scheme and must point to
// a resource that returns a valid feed. (ie: https:whatever.com/feed).
//
// GetFeed uses github.com/mmcdole/gofeed to make the request
// and parse the response body.
func (api *API) GetFeed(feedUrl string) (*gofeed.Feed, error) {
	feed := &gofeed.Feed{}
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(feedUrl)

	// fmt.Println("feed type: ", reflect.TypeOf(feed).String())

	if err != nil {
		fmt.Println("feed parsing err: ", err)
		return feed, err
	}

	return feed, nil
}

func (api *API) GetFeedsConcurrent(feedUrls []string) ([]*gofeed.Feed, error) {
	feeds := []*gofeed.Feed{}
	fp := gofeed.NewParser()
	type Result struct {
		Res   *gofeed.Feed
		Error error
	}

	ch := make(chan Result, len(feedUrls))

	for _, link := range feedUrls {
		u := link
		go func() {
			res, err := fp.ParseURL(u)
			if err != nil {
				ch <- Result{
					Res:   &gofeed.Feed{},
					Error: err,
				}
			} else {
				ch <- Result{
					Res:   res,
					Error: nil,
				}
			}
		}()
	}

	for i := 0; i < len(feedUrls); i++ {
		result := <-ch
		if result.Error == nil {
			feeds = append(feeds, result.Res)
		}
	}

	return feeds, nil
}

// GuessFeedLinks attempts to guess the endpoint
// where an RSS/Atom/JSON feed lives given a valid
// URL (ie: https://siteurl.com).
// Note, it should be called only after after api.FindFeedLinks
// has failed
func (api *API) GuessFeedLinks(searchUrl string) ([]string, error) {
	confirmed := []string{}
	guesses := []string{}

	validURL, err := util.MakeUrl(searchUrl)
	if err != nil {
		return confirmed, err
	}
	siteUrl := validURL.String()

	// for each common feed path
	// attempt to create a valid url
	// if valid, append to potentialGuesses
	for _, path := range common.CommonFeedPaths {
		u, err := url.ParseRequestURI(siteUrl + path)
		if err != nil {
			return confirmed, err
		}
		guesses = append(guesses, u.String())
	}

	// For each guess, append one of the common feed
	// extensions, as long as the guess does not already
	// end with "/" or one of the common extensions
	// for _, guess := range guesses {
	// 	for _, ext := range common.CommonFeedExtensions {
	// 		if !strings.HasSuffix(guess, "/") && !strings.HasSuffix(guess, ext) {
	// 			withExt := guess + ext
	// 			guesses = append(guesses, withExt)
	// 		}
	// 	}
	// }

	// Create len(guesses) channels
	// and make requests concurrently.
	// If we receive httpStatusOK &&
	// header.Content-Type is in ValidContentTypes
	// we probably have a match
	type Result struct {
		Res   *http.Response
		Error error
	}
	ch := make(chan Result, len(guesses))

	for _, guess := range guesses {
		u := guess
		go func() {
			res, err := api.Client.Get(u)
			if err != nil {
				ch <- Result{
					Res:   &http.Response{},
					Error: err,
				}
			}

			if res.StatusCode != http.StatusOK {
				ch <- Result{
					Res:   res,
					Error: fmt.Errorf("bad result"),
				}
			} else {
				ch <- Result{
					Res:   res,
					Error: nil,
				}
			}
		}()
	}

	for _, guess := range guesses {
		result := <-ch
		if result.Error == nil {
			res := result.Res

			if res.Body != nil {
				defer res.Body.Close()
				contentType := res.Header.Get("Content-Type")
				if slices.Contains(common.ValidDocContentTypes, common.DocContentType(contentType)) {
					confirmed = append(confirmed, guess)
				}
			}
		}
	}
	return confirmed, nil
}

// FindFeedLinks searches the document at a given URL for
// feed links. The scheme ("https:") of the URL can be omitted
// and this function will try to create a full url.URL by prepending
// "https:".
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

		if !slices.Contains(common.ValidContentTypes, common.ContentType(linkType)) {
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
