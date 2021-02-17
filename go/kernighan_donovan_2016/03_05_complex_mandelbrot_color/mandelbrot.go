// PNG image of the Mandelbrot fractal
package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
)

func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, 2, 2
        width, height = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)
            // img point (px, py) represents complex value of z
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
    // those are completely arbitrally values
    const (
        iterations = 0xff
        contrast = 0x00aff0
        alfa = 0xff
    )

    var v complex128
    black := color.RGBA{0x00, 0x00, 0x00, alfa}
    c := 0x000000
    redmask := 0xff0000
    greenmask := 0x00ff00
    bluemask := 0x0000ff
    var red, green, blue uint8
    for n := uint8(0); n < iterations; n++ {
        c += contrast
        v = v * v + z
        if cmplx.Abs(v) > 2 {
            red = uint8((c & redmask) >> 16)
            green = uint8((c & greenmask) >> 8)
            blue = uint8(c & bluemask)
            return color.RGBA{red, green, blue, uint8(alfa)}
        }
    }
    return black
}
