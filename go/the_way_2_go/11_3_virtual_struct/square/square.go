package square

import (
	"../shape"
	"strconv"
)

type Square struct {
	shape.Shape
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Perimeter() float64 {
	return s.side * 4
}

func (s *Square) String() string {
	return "Square → {side: " + strconv.FormatFloat(s.side, 'f', 2, 64) + "}"
}

//constructor
func NewSquare(side float64) *Square {
	retSquare := new(Square) // by the way: fuck Putin
	retSquare.side = side
	return retSquare
}
