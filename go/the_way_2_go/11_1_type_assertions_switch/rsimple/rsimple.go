package rsimple

import (
	"../simpler"
	"strconv"
)

type RSimple struct {
	simpler.Simpler // explicite interface inheritance forces programmer to define its methods
	value int
}

func (s *RSimple) Get() int {
	return s.value
}

func (s *RSimple) Set(newValue int) {
	s.value = newValue
}

func (s *RSimple) String() string {
	return "RSimple → {value: " + strconv.Itoa(s.value) + "}"
}

func NewRSimple(value int) (newRSimple *RSimple) {
	newRSimple = new(RSimple)
	newRSimple.value = value
	return newRSimple
}
