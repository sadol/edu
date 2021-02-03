package square

import (
	"../shaper"
	"strconv"
)

type Square struct {
	shaper.Shaper
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) String() string {
	return "Square → {side: " + strconv.FormatFloat(s.side, 'f', 2, 64) + "}"
}

func NewSquare(sid float64) *Square {
	var squareToReturn = new(Square)
	squareToReturn.side = sid
	return squareToReturn
}
