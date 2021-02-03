// helper package of point structs & point related functions

package point

import "math"

type Point struct {
	X, Y float64
}

type Point3D struct {
	X, Y, Z float64
}

type PointPolar struct {
	Radius, Angle float64
}

func Abs(point *Point) float64 {
	return math.Hypot(point.X, point.Y)
}

func Scale(point *Point, scale float64) *Point {
	point.X *= scale
	point.Y *= scale
	return point
}
