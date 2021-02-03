package rectangle

import (
    "fmt"
    "math"
)

const errorMsg = "Rectangle < %v > too large."

type Rectangle struct {
	Hight, Width float64
}

// struct `Rectange' method; calculates area of the rectangle
func (rec Rectangle) Area() (area float64, err error) {
    sideMax := 0.5 * math.MaxFloat64                   // better safe than sorry
    if rec.Hight >= sideMax || rec.Width >= sideMax {
        err = fmt.Errorf(errorMsg, rec)
    } else {
        area = rec.Hight * rec.Width
    }
	return
}

// struct `Rectange' method; calculates perimeter of the rectangle
func (rec Rectangle) Perimeter() (perimeter float64, err error) {
    sideMax := 0.25 * math.MaxFloat64                  // better safe than sorry
    if rec.Hight >= sideMax || rec.Width >= sideMax {
        err = fmt.Errorf(errorMsg, rec)
    } else {
        perimeter =  (rec.Hight * 2) + (rec.Width * 2)
    }
	return
}
