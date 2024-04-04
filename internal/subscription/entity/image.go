package entity

import (
	"encoding/json"
)

type Image struct {
	ID    int64
	URL   string
	Title string
}

func (i Image) JSON() []byte {
	json, _ := json.MarshalIndent(i, "", "  ")
	return json
}

func (i Image) String() string {
	return string(i.JSON())
}
