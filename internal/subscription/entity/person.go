package entity

import (
	"encoding/json"
)

type Person struct {
	ID    int64
	Name  string
	Email string
}

func (p Person) JSON() []byte {
	j, _ := json.MarshalIndent(p, "", "  ")
	return j
}

func (p Person) String() string {
	return string(p.JSON())
}
