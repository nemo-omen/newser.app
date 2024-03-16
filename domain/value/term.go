package value

type Term string

func NewTerm(term string) (Term, error) {
	if len(term) == 0 {
		return Term(""), ErrInvalidInput
	}
	return Term(term), nil
}

func (t Term) String() string {
	return string(t)
}
