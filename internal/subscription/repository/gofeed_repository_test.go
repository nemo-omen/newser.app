package repository

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewGofeedRepository(t *testing.T) {
	client := &http.Client{}
	repo := NewGofeedRepository(client)

	assert.Equal(t, 10*time.Second, client.Timeout)
	assert.NotNil(t, repo.Client)
}

func TestGofeedRepository_IsValidSite(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/valid" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	client := &http.Client{}
	repo := NewGofeedRepository(client)

	// Test valid site
	validSite := testServer.URL + "/valid"
	isValid := repo.IsValidSite(validSite)
	assert.True(t, isValid)

	// Test invalid site
	invalidSite := testServer.URL + "/invalid"
	isValid = repo.IsValidSite(invalidSite)
	assert.False(t, isValid)
}

func TestGofeedRepository_GetFeed(t *testing.T) {
	// Create a test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/valid" {
			w.WriteHeader(http.StatusOK)
			// Return a sample feed XML
			fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
				<rss version="2.0">
					<channel>
						<title>Test Feed</title>
						<link>http://example.com</link>
						<description>Test feed description</description>
						<item>
							<title>Test Item</title>
							<link>http://example.com/item</link>
							<description>Test item description</description>
						</item>
					</channel>
				</rss>`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	// Create a new GofeedRepository with a test client
	client := &http.Client{}
	repo := NewGofeedRepository(client)

	// Test valid feed
	validFeed := testServer.URL + "/valid"
	feed, err := repo.GetFeed(validFeed)
	assert.NoError(t, err)
	assert.NotNil(t, feed)
	assert.Equal(t, "Test Feed", feed.Title)
	assert.Equal(t, "Test feed description", feed.Description)
	assert.Equal(t, "http://example.com", feed.Link)
	assert.Equal(t, "http://example.com/item", feed.Items[0].Link)

	// Test invalid feed
	invalidFeed := testServer.URL + "/invalid"
	_, err = repo.GetFeed(invalidFeed)
	assert.Error(t, err)
	assert.EqualError(t, err, "GetFeed: fp.ParseURL: gofeed failed to parse feed")
}

func TestGofeedRepository_GetFeed_WithMissingFields(t *testing.T) {
	// Create a test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/valid" {
			w.WriteHeader(http.StatusOK)
			// Return a sample feed XML with missing fields
			fmt.Fprint(w, `
			<?xml version="1.0" encoding="UTF-8"?>
				<rss version="2.0">
					<channel>
						<title>Test Feed</title>
						<description>Test feed description</description>
						<pubDate>Wed, 14 Feb 2024 00:00:00 -0600</pubDate>

						<item>
							<title>Test Item</title>
							<description>Test item description</description>
						</item>
					</channel>
				</rss>
				`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	// Create a new GofeedRepository with a test client
	client := &http.Client{}
	repo := NewGofeedRepository(client)

	// Test feed with missing fields
	validFeed := testServer.URL + "/valid"
	feed, err := repo.GetFeed(validFeed)
	assert.NoError(t, err)
	assert.NotNil(t, feed)
	assert.Equal(t, "Test Feed", feed.Title)
	assert.Equal(t, "Test feed description", feed.Description) // Description should be empty
	assert.Equal(t, testServer.URL, feed.Link)                 // Link should be server URL
	assert.Equal(t, "Test Item", feed.Items[0].Title)
	assert.Equal(t, "Test item description", feed.Items[0].Description)
}

func TestGofeedRepository_GetFeeds(t *testing.T) {
	// Create a test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/valid" {
			w.WriteHeader(http.StatusOK)
			// Return a sample feed XML
			fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
				<rss version="2.0">
					<channel>
						<title>Test Feed</title>
						<link>http://example.com</link>
						<description>Test feed description</description>
						<item>
							<title>Test Item</title>
							<link>http://example.com/item</link>
							<description>Test item description</description>
						</item>
					</channel>
				</rss>`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	// Create a new GofeedRepository with a test client
	client := &http.Client{}
	repo := NewGofeedRepository(client)

	// Test GetFeeds
	feedUrls := []string{testServer.URL + "/valid", testServer.URL + "/invalid"}
	feeds, err := repo.GetFeeds(feedUrls)
	assert.NoError(t, err)
	assert.NotNil(t, feeds)
	assert.Len(t, feeds, 1)
	assert.Equal(t, "Test Feed", feeds[0].Title)
	assert.Equal(t, "Test feed description", feeds[0].Description)
	assert.Equal(t, "http://example.com", feeds[0].Link)
	assert.Equal(t, "http://example.com/item", feeds[0].Items[0].Link)
}

func TestGofeedRepository_GetFaviconSrc(t *testing.T) {
	// Create a test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/valid" {
			w.WriteHeader(http.StatusOK)
			// Return a sample HTML document with favicon link
			fmt.Fprint(w, `
				<html>
					<head>
						<link rel="icon" href="/favicon.ico">
					</head>
					<body></body>
				</html>
			`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	// Create a new GofeedRepository with a test client
	client := &http.Client{}
	repo := NewGofeedRepository(client)

	// Test GetFaviconSrc
	siteUrl := testServer.URL + "/valid"
	faviconSrc := repo.GetFaviconSrc(siteUrl)
	assert.Equal(t, "/favicon.ico", faviconSrc)
}
func TestGofeedRepository_FindFeedLinks(t *testing.T) {
	// Create a test server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/valid" {
			w.WriteHeader(http.StatusOK)
			// Return a sample HTML document with feed links
			fmt.Fprint(w, `
				<html>
					<head>
						<link rel="alternate" type="application/rss+xml" href="/feed1.xml">
						<link rel="alternate" type="application/atom+xml" href="/feed2.xml">
						<link rel="alternate" type="application/json" href="/feed3.json">
						<link rel="stylesheet" href="/styles.css">
					</head>
					<body></body>
				</html>
			`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	// Create a new GofeedRepository with a test client
	client := &http.Client{}
	repo := NewGofeedRepository(client)

	// Test FindFeedLinks
	siteUrl := testServer.URL + "/valid"
	links, err := repo.FindFeedLinks(siteUrl)
	assert.NoError(t, err)
	assert.NotNil(t, links)
	assert.Len(t, links, 3)
	assert.Contains(t, links, testServer.URL+"/feed1.xml")
	assert.Contains(t, links, testServer.URL+"/feed2.xml")
	assert.Contains(t, links, testServer.URL+"/feed3.json")
}

func TestGofeedRepository_GuessFeedLinks(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/atom" {
			w.Header().Add("Content-Type", "application/xml; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
				<feed xmlns="http://www.w3.org/2005/Atom">
					<title>Test Feed</title>
					<link>http://example.com</link>
					<subtitle>Test feed description</subtitle>
					<entry>
						<title>Test Item</title>
						<link>http://example.com/item</link>
						<summary>Test item description</summary>
					</entry>
				</feed>`)
		} else if r.URL.Path == "/rss" {
			w.Header().Add("Content-Type", "application/xml; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
				<rss version="2.0">
					<channel>
						<title>Test Feed</title>
						<link>http://example.com</link>
						<description>Test feed description</description>
						<item>
							<title>Test Item</title>
							<link>http://example.com/item</link>
							<description>Test item description</description>
						</item>
					</channel>
				</rss>`)
		} else if r.URL.Path == "/json" {
			w.Header().Add("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{
				"title": "Test Feed",
				"link": "http://example.com",
				"description": "Test feed description",
				"items": [
					{
						"title": "Test Item",
						"link": "http://example.com/item",
						"description": "Test item description"
					}
				]
			}`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	client := &http.Client{}
	repo := NewGofeedRepository(client)

	siteUrl := testServer.URL

	confirmed, err := repo.GuessFeedLinks(siteUrl)
	assert.NoError(t, err)
	assert.NotEmpty(t, confirmed)
}
