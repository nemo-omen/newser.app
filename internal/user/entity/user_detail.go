package entity

import (
	"encoding/json"
)

type UserDetail struct {
	ID    int64
	Name  string
	Email string
}

func (u UserDetail) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}

func (u UserDetail) String() string {
	return string(u.JSON())
}
