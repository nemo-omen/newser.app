package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"newser.app/internal/domain/value"
)

func TestNewsfeed(t *testing.T) {
	// Test data
	id := NewID()
	title := "Test Newsfeed"
	description := "This is a test newsfeed"
	feedLink := value.Link("https://example.com/feed")
	siteLink := value.Link("https://example.com")
	author := &Person{
		Name:  "John Doe",
		Email: "john@doe.whatever",
	}
	language := "en"
	image := &Image{
		URL:   "https://example.com/image.jpg",
		Title: "Example Corp.",
	}
	copyRight := "© 2022 Example Corp."
	articles := []*Article{
		{
			Item: &Item{
				ID:        NewID(),
				Title:     "Article 1",
				Content:   "This is article 1",
				Link:      value.Link("https://example.com/article1"),
				Updated:   "2022-01-01T00:00:00Z",
				Published: "2022-01-01T00:00:00Z",
				GUID:      "123456789",
			},
			Read:  false,
			Saved: false,
		},
		{
			Item: &Item{
				ID:        NewID(),
				Title:     "Article 2",
				Content:   "This is article 2",
				Link:      value.Link("https://example.com/article2"),
				Updated:   "2022-01-01T00:00:00Z",
				Published: "2022-01-01T00:00:00Z",
				GUID:      "987654321",
			},
			Read:  false,
			Saved: false,
		},
	}

	// Create a new newsfeed
	newsfeed := &Newsfeed{
		ID:          id,
		Title:       title,
		Description: description,
		FeedLink:    feedLink,
		SiteLink:    siteLink,
		Author:      author,
		Language:    language,
		Image:       image,
		Copyright:   copyRight,
		Articles:    articles,
	}

	// Assert the values
	assert.Equal(t, id, newsfeed.ID)
	assert.Equal(t, title, newsfeed.Title)
	assert.Equal(t, description, newsfeed.Description)
	assert.Equal(t, feedLink, newsfeed.FeedLink)
	assert.Equal(t, siteLink, newsfeed.SiteLink)
	assert.Equal(t, author, newsfeed.Author)
	assert.Equal(t, language, newsfeed.Language)
	assert.Equal(t, image, newsfeed.Image)
	assert.Equal(t, copyRight, newsfeed.Copyright)
	assert.Equal(t, articles, newsfeed.Articles)
}

func TestNewNewsfeed(t *testing.T) {
	// Test data
	title := "Test Newsfeed"
	description := "This is a test newsfeed"
	feedLink := "https://example.com/feed"
	siteLink := "https://example.com"
	language := "en"
	copyright := "© 2022 Example Corp."
	author := &Person{
		Name:  "John Doe",
		Email: "john@doe.doe",
	}
	image := &Image{
		URL:   "https://example.com/image.jpg",
		Title: "Example Corp.",
	}

	nf := NewNewsfeed(title, description, feedLink, siteLink, language, copyright, author, image)

	assert.NotNil(t, nf)
	assert.Equal(t, title, nf.Title)
	assert.Equal(t, description, nf.Description)
	assert.Equal(t, value.Link(feedLink), nf.FeedLink)
	assert.Equal(t, value.Link(siteLink), nf.SiteLink)
	assert.Equal(t, author, nf.Author)
	assert.Equal(t, language, nf.Language)
	assert.Equal(t, image, nf.Image)
}

func TestNewsfeed_AddArticle(t *testing.T) {
	// Test data
	nf := NewNewsfeed("Test Newsfeed", "This is a test newsfeed", "https://example.com/feed", "https://example.com", "en", "© 2022 Example Corp.", &Person{
		Name:  "John Doe",
		Email: "john@doe.doe",
	}, &Image{
		URL:   "https://example.com/image.jpg",
		Title: "Example Corp.",
	})

	article := &Article{
		Item: &Item{
			ID:        NewID(),
			Title:     "Article 1",
			Content:   "This is article 1",
			Link:      value.Link("https://example.com/article1"),
			Updated:   "2022-01-01T00:00:00Z",
			Published: "2022-01-01T00:00:00Z",
			GUID:      "123456789",
		},
		Read:  false,
		Saved: false,
	}

	// Add the article
	nf.AddArticle(article)

	// Assert the article was added
	assert.Equal(t, 1, len(nf.Articles))
	assert.Equal(t, article, nf.Articles[0])
}
