package repository

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"newser.app/pkg/constant"
)

// isValidSite tests whether a given url leads to
// a valid site by sending a HEAD request and
// returning whether the response StatusCode == 200
// If the request results in an error, the result is false.
func isValidSite(siteUrl string, client http.Client) bool {
	res, err := client.Head(siteUrl)
	if err != nil {
		return false
	}

	if res.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func getFaviconSrc(siteUrl string, client http.Client) string {
	src := ""
	res, err := client.Get(siteUrl)
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
		if !strings.Contains(rel, "icon") {
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

// guessFeedLinks attempts to guess the endpoint
// where an RSS/Atom/JSON feed lives given a valid
// URL (ie: https://siteurl.com).
// Note, it should be called only after after api.FindFeedLinks
// has failed
func guessFeedLinks(siteUrl string, client http.Client) ([]string, error) {
	confirmed := []string{}
	guesses := []string{}

	// for each common feed path
	// attempt to create a valid url
	// if valid, append to potentialGuesses
	for _, path := range constant.CommonFeedPaths {
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
			res, err := client.Get(u)
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
				if slices.Contains(constant.ValidDocContentTypes, constant.DocContentType(contentType)) {
					confirmed = append(confirmed, guess)
				}
			}
		}
	}
	return confirmed, nil
}
