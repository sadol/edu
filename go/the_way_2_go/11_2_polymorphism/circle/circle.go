package circle

import (
	"math"
	"strconv"
)

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c *Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func (c *Circle) String() string {
	return "Circle → {radius: " + strconv.FormatFloat(c.radius, 'f', 2, 64) + "}"
}

//constructor
func NewCircle(radius float64) *Circle {
	retCircle := new(Circle)
	retCircle.radius = radius
	return retCircle
}
