package point

import (
	"../checkFloat"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	// INFO: math.Hypot handles overflow
	return math.Hypot(p.X, p.Y)
}

func (p *Point) Scale(factor float64) bool {
	canBeDoneX, canBeDoneY := checkFloat.CheckMultOverflow(p.X, factor), checkFloat.CheckMultOverflow(p.Y, factor)

	if canBeDoneX && canBeDoneY {
		p.X *= factor
		p.Y *= factor
		return true
	}
	return false
}
