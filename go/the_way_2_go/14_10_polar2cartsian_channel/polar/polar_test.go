// tests of the polar coordinates structs and methods
package polar                                     // the same package as testee

import (
    "testing"
)

func TestGetCoordinates(t *testing.T) {
    r, phi := 1.000, -1.001
    somePolar := NewPolar(r, phi)
    gotR, gotPhi := somePolar.GetCoordinates()
    if gotR != r || gotPhi != phi {
        t.Errorf("GetCoordinates should return %g,%g but returned %g,%g instead.",
                 r, phi, gotR, gotPhi)
    }
}

/* DO NOT COMPARE FLOATS!!!
func TestPolarToCartesian(t *testing.T) {
    x, y := 0.000, 1.000
    r, phi := 1.000, math.Pi/2
    somePolar := NewPolar(r, phi)
    var gotCartesian *cartesian.Cartesian = somePolar.PolarToCartesian()
    gotX, gotY := gotCartesian.GetCoordinates()
    if gotX != x || gotY != y {
        t.Errorf("PolarToCartesian should return %g, %g but returned %g, %g instead.",
                 x, y, gotX, gotY)
    }
}
*/
