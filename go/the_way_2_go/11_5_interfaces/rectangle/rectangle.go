package rectangle

import (
	"../shaper"
	"strconv"
)

type Rectangle struct {
	shaper.Shaper
	width  float64
	heigth float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.heigth
}

func (r *Rectangle) Perimeter() float64 {
	return (r.width * 2) + (r.heigth * 2)
}

func (r *Rectangle) String() string {
	return "Rectanlge → {width: " + strconv.FormatFloat(r.width, 'f', 2, 64) +
		", heigth: " + strconv.FormatFloat(r.heigth, 'f', 2, 64) + "}"
}

func NewRectangle(wid, hei float64) *Rectangle {
	var rectangleToReturn = new(Rectangle)
	rectangleToReturn.width = wid
	rectangleToReturn.heigth = hei
	return rectangleToReturn
}
