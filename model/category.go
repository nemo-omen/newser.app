package model

type Category struct {
	Term      string
	Articles  []*Article
	Newsfeeds []*Newsfeed
}
