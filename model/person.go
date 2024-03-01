package model

type Person struct {
	ID       int64
	Name     string
	Email    string
	Articles []Article
}
