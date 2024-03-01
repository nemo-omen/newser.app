package model

type Subscription struct {
	Id         int64 `db:"id"`
	NewsfeedId int64 `db:"newsfeed_id"`
	UserId     int64 `db:"user_id"`
}
