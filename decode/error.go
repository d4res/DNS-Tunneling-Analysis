package decode

type ParseError struct{}

func (ParseError) Error() string {
	return "parse error"
}
