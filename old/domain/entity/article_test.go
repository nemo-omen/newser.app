package entity

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"newser.app/internal/domain/value"
)

func TestNewArticle(t *testing.T) {
	author := &Person{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	image := &Image{
		URL:   "https://example.com/image.jpg",
		Title: "Example Image",
	}

	categories := []string{"category1", "category2"}

	article := NewArticle(
		"Example Title",
		"Example Description",
		"Example Content",
		"https://example.com",
		"2022-01-01",
		"2022-01-01",
		"example-guid",
		time.Date(2022, 01, 01, 00, 00, 0, 0, time.UTC),
		time.Date(2022, 01, 01, 00, 00, 0, 0, time.UTC),
		author,
		image,
		categories,
	)

	assert.NotNil(t, article)
	assert.Equal(t, "Example Title", article.Title)
	assert.Equal(t, "Example Description", article.Description)
	assert.Equal(t, "Example Content", article.Content)
	assert.Equal(t, "https://example.com", article.Link.String())
	assert.Equal(t, "2022-01-01", article.Updated)
	assert.Equal(t, "2022-01-01", article.Published)
	assert.Equal(t, author, article.Author)
	assert.Equal(t, "example-guid", article.GUID)
	assert.Equal(t, image, article.Image)
	assert.Len(t, article.Categories, 2)
	assert.Equal(t, value.Term("category1"), article.Categories[0].Term)
	assert.Equal(t, value.Term("category2"), article.Categories[1].Term)
}

func TestArticle_JSON(t *testing.T) {
	author := &Person{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	image := &Image{
		URL:   "https://example.com/image.jpg",
		Title: "Example Image",
	}

	categories := []string{"category1", "category2"}

	article := NewArticle(
		"Example Title",
		"Example Description",
		"Example Content",
		"https://example.com",
		"2022-01-01",
		"2022-01-01",
		"example-guid",
		time.Date(2022, 01, 01, 00, 00, 0, 0, time.UTC),
		time.Date(2022, 01, 01, 00, 00, 0, 0, time.UTC),
		author,
		image,
		categories,
	)

	jsonData := article.JSON()

	var decodedArticle Article
	err := json.Unmarshal(jsonData, &decodedArticle)
	assert.NoError(t, err)

	assert.Equal(t, article.ID, decodedArticle.ID)
	assert.Equal(t, article.Title, decodedArticle.Title)
	assert.Equal(t, article.Description, decodedArticle.Description)
	assert.Equal(t, article.Content, decodedArticle.Content)
	assert.Equal(t, article.Link.String(), decodedArticle.Link.String())
	assert.Equal(t, article.Updated, decodedArticle.Updated)
	assert.Equal(t, article.Published, decodedArticle.Published)
	assert.Equal(t, article.Author, decodedArticle.Author)
	assert.Equal(t, article.GUID, decodedArticle.GUID)
	assert.Equal(t, article.Image, decodedArticle.Image)
	assert.Len(t, decodedArticle.Categories, 2)
	assert.Equal(t, value.Term("category1"), decodedArticle.Categories[0].Term)
	assert.Equal(t, value.Term("category2"), decodedArticle.Categories[1].Term)
}

func TestArticle_SetRead(t *testing.T) {
	article := &Article{}
	assert.False(t, article.Read)

	article.SetRead()
	assert.True(t, article.Read)
}

func TestArticle_SetSaved(t *testing.T) {
	article := &Article{}
	assert.False(t, article.Saved)

	article.SetSaved()
	assert.True(t, article.Saved)
}

func TestArticle_SetUnread(t *testing.T) {
	article := &Article{Read: true}
	assert.True(t, article.Read)

	article.SetUnread()
	assert.False(t, article.Read)
}
