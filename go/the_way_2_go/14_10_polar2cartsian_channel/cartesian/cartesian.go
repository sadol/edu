// 2D point structure (Cartesian coordination version)
package cartesian

import (
    //"fmt"
    "strconv"
    //"../polar"
    //"math"
)

// basic 2D point struct (Cartesian coordinates version)
type Cartesian struct {
    x, y float64
}

// getter
func (c *Cartesian) GetCoordinates() (x, y float64) {
    return c.x, c.y
}

// setter
func (c *Cartesian) SetCoordinates(x, y float64) {
    c.x = x
    c.y = y
}

// stringificator of the `Cartesian struct'
func (c *Cartesian) String() (output string) {
    return "2D point(cartesian coordiantes) → {x: " +
           strconv.FormatFloat(c.x, 'f', 2, 64) + ", y:" +
           strconv.FormatFloat(c.y, 'f', 2, 64) + "}"
}
/* ---> to avoid circular import
// converts 2D point coordinates from Cartesian to polar
func (c *Cartesian) CartesianToPolar() (output *polar.Polar) {
    r := math.Sqrt((c.x * c.x) + (c.y * c.y))
    phi := math.Atan2(c.y / c.x)
    output = polar.NewPolar(r, phi)
    return
}
*/
// `Cartesian struct' constructor
func NewCartesian(x, y float64) (output *Cartesian) {
    output = new(Cartesian)
    output.x = x
    output.y = y
    return output
}
