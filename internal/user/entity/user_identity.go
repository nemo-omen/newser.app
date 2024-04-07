package entity

import (
	"encoding/json"
)

type UserIdentity struct {
	ID    int64
	Email string
}

func (u UserIdentity) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}

func (u UserIdentity) String() string {
	return string(u.JSON())
}
