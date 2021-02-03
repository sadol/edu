package simple

import "strconv"

type Simple struct {
	value int
}

func (s *Simple) Get() int {
	return s.value
}

func (s *Simple) Set(newValue int) {
	s.value = newValue
}

func (s *Simple) String() string {
	return "Simple → {value: " + strconv.Itoa(s.value) + "}"
}

func NewSimple(value int) *Simple {
	newSimple := new(Simple)
	newSimple.value = value
	return newSimple
}
