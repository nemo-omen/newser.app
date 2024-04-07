package entity

import "encoding/json"

type UserCreate struct {
	Name  string
	Email string
}

func (u UserCreate) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}

func (u UserCreate) String() string {
	return string(u.JSON())
}
