// package creates structs and methods for polar coordinates if the 2D point
package polar

import (
    //"fmt"
    "strconv"
    "math"
    "../cartesian"
)

// basic 2D point struct (polar coordinates version)
type Polar struct {
    r, phi float64
}

// getter
func (p *Polar) GetCoordinates() (r, phi float64) {
    return p.r, p.phi
}

// setter
func (p *Polar) SetCoordinates(r, phi float64) {
    p.r = r
    p.phi = phi
}

// stringificator of the `Polar struct'
func (p *Polar) String() (output string) {
    return "2D point(polar coordiantes) → {r: " +
           strconv.FormatFloat(p.r, 'f', 2, 64) + ", phi:" +
           strconv.FormatFloat(p.phi, 'f', 2, 64) + "}"
}

// converts 2D point coordinates from polar to Cartesian
func (p *Polar) PolarToCartesian() (output *cartesian.Cartesian) {
    x := p.r * math.Cos(p.phi)
    y := p.r * math.Sin(p.phi)
    output = cartesian.NewCartesian(x, y)
    return
}

// `Polar struct' constructor
func NewPolar(r, phi float64) (output *Polar) {
    output = new(Polar)
    output.r = r
    output.phi = phi
    return output
}
