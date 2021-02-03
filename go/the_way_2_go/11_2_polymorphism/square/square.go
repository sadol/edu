package square

import "strconv"

type Square struct {
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
func NewSquare(side float64) (retSquare *Square) {
	retSquare := new(Square)
	retSquare.side = side
	return retSquare
}
