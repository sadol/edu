// PNG image of the Mandelbrot fractal, server version with zooming
// functionality.

package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "net/http"
    "log"
    "strconv"
)

const (
    xminDefault, yminDefault, xmaxDefault, ymaxDefault = -2.00, -2.00, 2.00, 2.00
    xDefault, yDefault, zDefault = 0.00, 0.00, 1.00
    width, height = 1024, 1024
)

var (
    err error
    x, y, z float64
    xString, yString, zString string
)

func main() {
    http.HandleFunc("/", mandelHand)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelHand(w http.ResponseWriter, r *http.Request) {
    // ------------------x,y,zoom handling with URL strings---------------
    xmin, xmax, ymin, ymax := xminDefault, xmaxDefault, yminDefault, ymaxDefault
    if err = r.ParseForm(); err != nil {
        log.Print(err)
    }
    xString = r.Form.Get("x")
    if xString != "" {
        if x, err = strconv.ParseFloat(xString, 64); err != nil {
            log.Print(err)
        }
    } else {
        x = xDefault
    }
    yString = r.Form.Get("y")
    if yString != "" {
        if y, err = strconv.ParseFloat(yString, 64); err != nil {
            log.Print(err)
        }
    } else {
        y = yDefault
    }
    zString := r.Form.Get("z")
    if zString != "" {
        if z, err = strconv.ParseFloat(zString, 64); err != nil {
            log.Print(err)
        }
        // zooming
        xmin = xminDefault / z
        xmax = xmaxDefault / z
        ymin = yminDefault / z
        ymax = ymaxDefault / z
    } else {
        z = zDefault
    }

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        yy := (float64(py) / height * (ymax - ymin) + ymin) + y
        for px := 0; px < width; px++ {
            xx := (float64(px) / width * (xmax - xmin) + xmin) + x
            zz := complex(xx, yy)
            // img point (px, py) represents complex value of z
            img.Set(px, py, mandelbrot(zz))
        }
    }
    png.Encode(w, img)
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
