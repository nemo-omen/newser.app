package entity

import "encoding/json"

type Person struct {
	Name  string
	Email string
}

// JSON returns the JSON encoding of the Person.
func (p *Person) JSON() []byte {
	j, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return []byte{}
	}
	return j
}

// String returns the string representation of the Person.
func (p *Person) String() string {
	return string(p.JSON())
}
