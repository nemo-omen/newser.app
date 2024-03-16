package value

type Name string

func NewName(str string) (Name, error) {
	if len(str) == 0 {
		return Name(""), ErrInvalidInput
	}
	return Name(str), nil
}

func (n Name) String() string {
	return string(n)
}
