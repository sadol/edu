package main

import (
	"./circle"
	"./rectangle"
	"./shaper"
	"./square"
	"./triangle"
	"fmt"
)

func main() {
    // these are pointers
	var c1 = circle.NewCircle(2.22)
	var c2 = circle.NewCircle(2.23)
	var c3 = circle.NewCircle(12.23)
	var t1 = triangle.NewTriangle(2, 3)
	var t2 = triangle.NewTriangle(12, 3)
	var t3 = triangle.NewTriangle(2, 13)
	var r1 = rectangle.NewRectangle(2, 4)
	var r2 = rectangle.NewRectangle(12, 4)
	var r3 = rectangle.NewRectangle(2, 14)
	var s1 = square.NewSquare(2)
	var s2 = square.NewSquare(2.03)
	var s3 = square.NewSquare(13.44444)
    // and this is array of interfaces
	var myShapes = []shaper.Shaper{c1, c2, c3, t1, t2, t3, r1, r2, r3, s1, s2, s3}

	fmt.Println("All my shapes are here:")
	for _, singleShape := range myShapes {
		fmt.Printf("\t%v.\n", singleShape)
	}
	fmt.Println()
	fmt.Println("`PeriInterface' in use on table of shapes:")
	// new interface for easy cycling through shapes tables
	for _, singleShape := range myShapes {
		fmt.Printf("\tType: %T, perimeter: %f.\n", singleShape, singleShape.Perimeter())
	}
}
