package rectangle

import (
	"../shape"
	"strconv"
)

type Rectangle struct {
	shape.Shape
	length float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return (r.length * 2) + (r.width * 2)
}

func (r *Rectangle) String() string {
	return "Rectangle → {length: " + strconv.FormatFloat(r.length, 'f', 2, 64) +
		", width: " + strconv.FormatFloat(r.width, 'f', 2, 64) + "}"
}

//constructor
func NewRectangle(length float64, width float64) *Rectangle {
	retRectangle := new(Rectangle)
	retRectangle.length = length
	retRectangle.width = width
	return retRectangle
}
