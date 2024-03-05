package service

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/PuerkitoBio/goquery"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
	"newser.app/shared"
)

type API struct {
	Client *http.Client
}

func NewAPI(client *http.Client) API {
	client.Timeout = 10 * time.Second
	return API{
		Client: client,
	}
}

// CheckSite tests whether a given url leads to
// a valid site by sending a HEAD request and
// returning whether the response StatusCode == 200
// If the request results in an error, the result is false.
func (api API) CheckSite(siteUrl string) bool {
	// fmt.Printf("Checking whether %v is a valid site", siteUrl)
	res, err := api.Client.Head(siteUrl)
	if err != nil {
		return false
	}

	if res.StatusCode != http.StatusOK {
		return false
	}

	// fmt.Println(res.Status)
	return true
}

// GetFeed attempts to retrieve a valid RSS/Atom/JSON feed
// from the given URL string.
//
// The URL string must include a scheme and must point to
// a resource that returns a valid feed. (ie: https:whatever.com/feed).
//
// GetFeed uses github.com/mmcdole/gofeed to make the request
// and parse the response body.
func (api API) GetFeed(feedUrl string) (*gofeed.Feed, error) {
	feed := &gofeed.Feed{}
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(feedUrl)

	if err != nil {
		// fmt.Println("feed parsing err: ", err)
		return feed, err
	}

	// ensure feed description is free of html tags
	feed.Description = strip.StripTags(feed.Description)
	// fmt.Println("description: ", feed.Description)

	// double-check for site favicon if feed.Image
	// is not present
	if feed.Image == nil {
		u, _ := url.Parse(feedUrl)
		if u != nil {
			if u.Path != "" {
				u.Path = "/"
			}

			link := u.String()
			src := api.GetFaviconSrc(link)

			if src != "" {
				srcUrl, _ := url.Parse(src)
				if srcUrl != nil {
					if srcUrl.Scheme == "" {
						srcUrl.Scheme = u.Scheme
					}

					if srcUrl.Host == "" {
						srcUrl.Host = u.Host
					}
					src = srcUrl.String()
				}

				feed.Image = &gofeed.Image{
					URL:   src,
					Title: feed.Title,
				}
			}
		}
	}

	fmt.Printf("image: %+v\n", feed.Image)

	for _, item := range feed.Items {
		// strip and truncate item description
		item.Description = strip.StripTags(item.Description)
		if len(item.Description) > 256 {
			item.Description = item.Description[0:256] + "..."
		}
	}

	return feed, nil
}

func (api API) GetFeedsConcurrent(feedUrls []string) ([]*gofeed.Feed, error) {
	feeds := []*gofeed.Feed{}
	// fp := gofeed.NewParser()
	type Result struct {
		Res   *gofeed.Feed
		Error error
	}

	ch := make(chan Result, len(feedUrls))

	for _, link := range feedUrls {
		// fmt.Println("Attempting to retrieve feed from ", link)
		u := link
		go func() {
			res, err := api.GetFeed(u)
			if err != nil {
				ch <- Result{
					Res:   nil,
					Error: err,
				}
			} else {
				// substitute feed's defined
				// link for valid url we just got
				// the feed from. Why? Because
				// sometimes people forget to update
				// the url in their RSS feeds. If we've
				// found a valid feed at this URL, it
				// obviously works, so we may as well use it.
				res.FeedLink = u
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
func (api API) GuessFeedLinks(siteUrl string) ([]string, error) {
	// fmt.Println("Attempting to guess links at ", siteUrl)
	confirmed := []string{}
	guesses := []string{}

	// for each common feed path
	// attempt to create a valid url
	// if valid, append to potentialGuesses
	for _, path := range shared.CommonFeedPaths {
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
					Res:   nil,
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
				if slices.Contains(shared.ValidDocContentTypes, shared.DocContentType(contentType)) {
					confirmed = append(confirmed, guess)
				}
			}
		}
	}
	return confirmed, nil
}

func (api *API) GetFaviconSrc(siteUrl string) string {
	src := ""
	res, err := api.Client.Get(siteUrl)
	if err != nil {
		return src
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return src
	}

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return src
	}

	document.Find("link").Each(func(i int, el *goquery.Selection) {
		rel, exists := el.Attr("rel")
		if !exists {
			return
		}
		if rel != "shortcut icon" {
			return
		}

		href, exists := el.Attr("href")

		if !exists {
			return
		}

		src = href
	})
	return src
}

// FindFeedLinks searches the document at a given URL for
// feed links.
// siteUrl should be a valid URL (ie: https://whatever.com)
func (api API) FindFeedLinks(siteUrl string) ([]string, error) {
	// fmt.Println("Attempting to find feed links at", siteUrl)
	links := []string{}

	fullURL, err := url.ParseRequestURI(siteUrl)
	if err != nil {
		return links, err
	}

	res, err := api.Client.Get(siteUrl)
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

		if !slices.Contains(shared.ValidContentTypes, shared.ContentType(linkType)) {
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
			testUrl.Scheme = fullURL.Scheme
			testUrl.Host = fullURL.Host
			links = append(links, testUrl.String())
		} else {
			links = append(links, href)
		}
	})

	return links, nil
}
