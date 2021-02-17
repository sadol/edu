// PNG image of the Mandelbrot fractal with supersampling
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
        superFactor = 2                   // every pixel contains 4 "subpixels"
    )

    // bigger picture
    imgSuper := image.NewRGBA(image.Rect(0, 0, width * superFactor,
                              height * superFactor))
    for py := 0; py < height * superFactor; py++ {
        y:= float64(py) / (height * superFactor) * (ymax - ymin) + ymin
        for px := 0; px < width * superFactor; px++ {
            x := float64(px) / (width * superFactor) * (xmax - xmin) + xmin
            // img point (px, py) represents complex value of z
            imgSuper.Set(px, py, mandelbrot(complex(x, y)))
        }
    }

    // smaller picture
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height * superFactor; py++ {
        for px := 0; px < width * superFactor; px++ {
            red1, green1, blue1, alfa := imgSuper.RGBAAt(int(px), int(py)).RGBA()
            red2, green2, blue2, _ := imgSuper.RGBAAt(int(px + 1), int(py)).RGBA()
            red3, green3, blue3, _ := imgSuper.RGBAAt(int(px), int(py + 1)).RGBA()
            red4, green4, blue4, _ := imgSuper.RGBAAt(int(px + 1), int(py + 1)).RGBA()
            red := (red1 + red2 + red3 + red4) / 4
            green := (green1 + green2 + green3 + green4) / 4
            blue := (blue1 + blue2 + blue3 + blue4) / 4
            colorAvg := color.RGBA{uint8(red), uint8(green), uint8(blue), uint8(alfa)}
            // img point (px, py) represents complex value of z
            img.Set(px / superFactor, py / superFactor, colorAvg)
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
