package main

import (
	"./point"
	"./point3D"
	"fmt"
)

func main() {
	myPoint := point.Point{X: 3, Y: 4}
	fmt.Println(myPoint)
	scaleFactor := float64(6.004)
	fmt.Printf("Length of the origin vector based on myPoint: %f.\n", myPoint.Abs())
	if myPoint.Scale(scaleFactor) {
		fmt.Println(myPoint)
	} else { // there should be specialized interface for error handling somewhere
		fmt.Println("Scale error")
	}

	myPoint3D := point3D.Point3D{point.Point{X: 3, Y: 4}, 5}
	fmt.Println(myPoint3D)
	if value, ok := myPoint3D.Abs(); ok {
		fmt.Printf("Length of the origin vector based on myPoint3D: %f.\n", value)
	} else {
		fmt.Println("Scale error")
	}
	if myPoint3D.Scale(scaleFactor) {
		fmt.Println(myPoint3D)
	} else { // there should be specialized interface for error handling somewhere
		fmt.Println("Scale error")
	}
	fmt.Printf("Point3d X: %f.\n", myPoint3D.X)
	fmt.Printf("Point3d Y: %f.\n", myPoint3D.Y)
	fmt.Printf("Point3d Z: %f.\n", myPoint3D.Z)
}
