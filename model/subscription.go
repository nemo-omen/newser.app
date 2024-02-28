package model

type Subscription struct {
	Id       int64
	Newsfeed *Newsfeed
	User     *User
	Slug     string
}
