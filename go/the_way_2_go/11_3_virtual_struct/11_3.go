package main

import (
	"./circle"
	"./rectangle"
	"./shaper"
	"./square"
	"fmt"
)

func main() {
	myRadius, mySide, myLength, myWidth := 23.4, 12.5, 23.5, 0.12
	mySquarePn := square.NewSquare(mySide)
	myRectanglePn := rectangle.NewRectangle(myLength, myWidth)
	myCirclePn := circle.NewCircle(myRadius)
	shapers := []shaper.Shaper{mySquarePn, myRectanglePn, myCirclePn}
	fmt.Println("Testing the shapers.")
	for i := range shapers {
		fmt.Printf("Shaper prints => %v, area: %4.2f, perimeter: %4.2f.\n",
			shapers[i], shapers[i].Area(), shapers[i].Perimeter())
	}
}
