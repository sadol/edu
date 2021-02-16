// comutes svg rendering 3D surface function
package main

import (
    "fmt"
    "math"
)

const (
    width, height = 600, 320                           // canvas size in pixels
    cells         = 100                                 // number of grid cells
    xyrange       = 30.0                    // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange               // pixels per x or y unit
    zscale        = height * 0.4                       //     pixels per z unit
    angle         = math.Pi / 6                   // angle of x , y axes (=30°)
    afactor       = 0.5                  // `a' factor for the egg-box function
    bfactor       = 2.0                  // `b' factor for the egg-box function
    X2            = 1.0                  // `x²' factor for the saddle function
    Y2            = 2.0                  // `y²' factor for the saddle function
    XY            = 3.0                  // `xy' factor for the saddle function
    X             = 4.0                  // `x'  factor for the saddle function
    Y             = 5.0                  // `y'  factor for the saddle function
    A             = 6.0                   // `a' factor for the saddle function
)

var (
    sin30 = math.Sin(angle)
    cos30 = math.Cos(angle)
)

func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
    "style='stroke: grey; fill: white; stroke-width: 0.7' " +
    "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i + 1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j + 1)
            dx, dy := corner(i + 1, j + 1)
            if math.IsInf(ay, 0) || math.IsInf(by, 0) || math.IsInf(cy, 0) || math.IsInf(dy, 0) {
                continue
            }
            fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
                        ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
    // find point (x, y) at the corner of the cell (i, j)
    x := xyrange * (float64(i) / cells - .5)
    y := xyrange * (float64(j) / cells - .5)

    // compute surface hight z
    z := f(x, y)                 // HERE you can put eggbox or saddle functions

    // project (x, y, z) isometrically into 2d svg canvas (sx, sy)
    sx := width / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
    return sx, sy
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

// additional functions:
func eggBox(x, y float64) float64 {
    return afactor * ((math.Sin(x / bfactor) + math.Sin(y / bfactor))
}

func saddle(x, y float64) float64 {
    return  (X2 * math.Sqrt(x)) + (Y2 * math.Sqrt(y)) + (XY *x * y) + x + y + A
}
