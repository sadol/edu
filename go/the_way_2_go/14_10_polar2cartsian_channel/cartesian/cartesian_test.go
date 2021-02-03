// tests of the cartersian coordinates structs and methods
package cartesian                                 // the same package as testee

import (
    "testing"
    //"../polar"
)

func TestGetCoordinates(t *testing.T) {
    x, y := 1.000, -1.001
    someCartesian := NewCartesian(x, y)
    gotX, gotY := someCartesian.GetCoordinates()
    if gotX != x || gotY != y {
        t.Errorf("GetCoordinates should return %g,%g but returned %g,%g instead.",
                 x, y, gotX, gotY)
    }
}
/*
func TestCartesianToPolar(t *testing.T) {
    x, y := 12.000, 5.000
    r, phi := 13.000, 22.600
    someCartesian := NewCartesian(x, y)
    gotPolar := someCartesian.CartesianToPolar()
    gotR, gotPhi := gotPolar.GetCoordiantes()
    if gotR != r || gotPhi != phi {
        t.Errorf("CartToPolar should return %g,%g but returned %g,%g instead.",
                 r, phi, gotR, gotPhi)
    }
}
*/
