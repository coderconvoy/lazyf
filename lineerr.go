package lazyf

import "fmt"

type LineErr struct {
	s string
	l int
}

type Liner interface {
	error
	Line() int
}

func NewLineErr(n int, s string) LineErr {
	return LineErr{
		s: s,
		l: n,
	}
}

func (l LineErr) Error() string {
	return fmt.Sprintf("Error on line %d: %s", l.l, l.s)
}

func (l LineErr) Line() int {
	return l.l
}
