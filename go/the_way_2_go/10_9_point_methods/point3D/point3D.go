package point3D

import (
	"../array"
	"../checkFloat"
	"../point"
	"math"
)

type Point3D struct {
	point.Point "point 2D is embedded and ready to use by referencting (example: p.Point.Scale())"
	Z           float64 "third dimention"
}

func (p *Point3D) Abs() (vectorLen float64, ok bool) {
	if p.check3DHypot() {
		vectorLen, ok = math.Sqrt((p.X*p.X)+(p.Y*p.Y)+(p.Z*p.Z)), true
	} else {
		vectorLen, ok = 0, false
	}
	return
}

//UNEXPORTED
//helper antioverflow function for func=sqrt(x²+y²+z²)→func=(ð²+ð²+ð²)
// where ð=max(x, y, z)
func (p *Point3D) check3DHypot() bool {
	_, bigest, _ := array.MinMaxArrFloat64(p.X, p.Y, p.Z)
	// (ð²+ð²+ð²)=3ð²=sqrt(3)ð*sqrt(3)ð
	if checkFloat.CheckMultOverflow(bigest*(math.Sqrt(3)), bigest*(math.Sqrt(3))) {
		return true
	}
	return false
}

func (p *Point3D) Scale(factor float64) bool {
	canBeDoneX := checkFloat.CheckMultOverflow(p.X, factor)
	canBeDoneY := checkFloat.CheckMultOverflow(p.Y, factor)
	canBeDoneZ := checkFloat.CheckMultOverflow(p.Z, factor)

	if canBeDoneX && canBeDoneY && canBeDoneZ {
		p.X *= factor
		p.Y *= factor
		p.Z *= factor
		return true
	}
	return false
}
