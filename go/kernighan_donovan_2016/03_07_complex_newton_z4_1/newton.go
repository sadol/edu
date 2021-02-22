// PNG image of the Mandelbrot fractal
package main

import (
    "image"
    "image/color"
    "image/png"
    //"math/cmplx"
    "os"
//    "math"
)

func main() {
    const (
        xmin, ymin, xmax, ymax = -1.5, -1.5, 1.5, 1.5
        width, height = 4 * 1024, 4 * 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)
            // img point (px, py) represents complex value of z
            img.Set(px, py, polynomial(z))
        }
    }
    png.Encode(os.Stdout, img)
}

func polynomial(z complex128) color.Color {
    // those are completely arbitrally values
    const (
        iterations = 0xff
        contrast = 0x010101
        alfa = 0xff
        delta = 0.01
    )

    var (
        black = color.RGBA{0x00, 0x00, 0x00, alfa}
        red = color.RGBA{0xff, 0x00, 0x00, alfa}
        green = color.RGBA{0x00, 0xff, 0x00, alfa}
        blue = color.RGBA{0x00, 0x00, 0xff, alfa}
        other = color.RGBA{0xaa, 0xaa, 0xaa, alfa}
    )

    root1 := complex(-1, 0)
    root2 := complex(1, 0)
    root3 := complex(0, -1)
    root4 := complex(0, 1)
    N := ((3 * z * z * z * z) - 1) / (4 * z * z * z)

    for n := uint8(0); n < iterations; n++ {
        if dist(N, root1)  < delta { // first root basin coloring
            return red
        }
        if dist(N, root2) < delta { // second root basin coloring
            return green
        }
        if dist(N, root3) < delta { // third root basin coloring
            return blue
        }
        if dist(N, root4) < delta { // forth root basin coloring
            return other
        }
        N = ((3 * N * N * N * N) - 1) / (4 * N * N * N)
    }
    return black
}

// distance between two complex numbers
func dist (s, v complex128) (ret float64) {
    diff := s - v
    ret = (real(diff) * real(diff)) + (imag(diff) * imag(diff))
    return
}
