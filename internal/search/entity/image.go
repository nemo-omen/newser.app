package entity

import "encoding/json"

type Image struct {
	URL   string
	Title string
}

func (i *Image) JSON() []byte {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return []byte{}
	}
	return j
}

func (i *Image) String() string {
	return string(i.JSON())
}
