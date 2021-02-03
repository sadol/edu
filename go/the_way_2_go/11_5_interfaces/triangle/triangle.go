package triangle

import (
	"../shaper"
	"math"
	"strconv"
)

type Triangle struct {
	shaper.Shaper
	base   float64
	heigth float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.heigth
}

func (t *Triangle) Perimeter() float64 {
	return t.base + t.heigth + math.Sqrt(math.Pow(t.base, 2)+math.Pow(t.heigth, 2))
}

func (t *Triangle) String() string {
	return "Trianlge → {base: " + strconv.FormatFloat(t.base, 'f', 2, 64) +
		", heigth: " + strconv.FormatFloat(t.heigth, 'f', 2, 64) + "}"
}

func NewTriangle(bas, hei float64) *Triangle {
	var triangleToReturn = new(Triangle)
	triangleToReturn.base = bas
	triangleToReturn.heigth = hei
	return triangleToReturn
}
