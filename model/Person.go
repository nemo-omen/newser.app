package model

import "encoding/json"

type Person struct {
	ID        int64       `json:"-,omitempty" db:"id"`
	Name      string      `json:"name,omitempty" db:"name"`
	Email     string      `json:"email,omitempty" db:"email"`
	Articles  []*Article  `json:"-" db:"-"`
	Newsfeeds []*Newsfeed `json:"-" db:"-"`
}

func (p Person) String() string {
	j, _ := json.MarshalIndent(p, "", "    ")
	return string(j)
}
