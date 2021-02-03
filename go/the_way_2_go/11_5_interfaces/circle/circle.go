package circle

import (
	"../shaper"
	"math"
	"strconv"
)

type Circle struct {
	shaper.Shaper
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) String() string {
	return "Circle → {radius: " + strconv.FormatFloat(c.radius, 'f', 2, 64) + "}"
}

func NewCircle(rad float64) *Circle {
	var circleToReturn = new(Circle)
	circleToReturn.radius = rad
	return circleToReturn
}
