package model

type Subscription struct {
	Id       uint
	Newsfeed *Newsfeed
	User     *User
	Slug     string
}
