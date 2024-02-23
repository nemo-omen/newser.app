package model

type Avatar struct {
	Id       uint
	Filename string
	Src      string
	Alt      string
}

type Email struct {
	Address   string
	Confirmed bool
}
