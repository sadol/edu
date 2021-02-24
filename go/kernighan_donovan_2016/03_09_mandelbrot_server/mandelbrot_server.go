// PNG image of the Mandelbrot fractal, server version with zooming
// functionality.
// WARNING: you can easly burn down your computer runnig this code
// INFO: builtin types version of this program runs very fast and is resonably
// accurate up to 100x-130x zoom but in 200x zoom low accuracy graphical artifacts
// are clearly visible.
package main

import (
    "image"
    "image/color"
    "image/png"
    "net/http"
    "log"
    "strconv"
    "math/big"
)

// -----------------------custom complex number type---------------------------
type MyComplex struct {
    Imag *big.Float
    Real *big.Float
}

// factory function for creating custom version of complex number
func NewMyComplex(r, i float64) (output MyComplex) {
    output = *(new(MyComplex))
    output.Imag = big.NewFloat(r)
    output.Real = big.NewFloat(i)
    return output
}

func (output MyComplex) Add (first, second MyComplex) MyComplex {
    output.Imag.Add(first.Imag, second.Imag)
    output.Real.Add(first.Real, second.Real)
    return output
}

func (output MyComplex) Mul(first, second MyComplex) MyComplex {
    output.Imag.Mul(first.Imag, second.Imag)
    output.Real.Mul(first.Real, second.Real)
    return output
}

func (first MyComplex) Abs () *big.Float {
    output1 := big.NewFloat(0)
    output2 := big.NewFloat(0)
    output1.Mul(first.Real, first.Real)
    output2.Mul(first.Imag, first.Imag)
    output1.Add(output1, output2)
    output1.Sqrt(output1)
    return output1
}
//-----------------------------------------------------------------------------

const (
    xminDefault, yminDefault, xmaxDefault, ymaxDefault = -2.00, -2.00, 2.00, 2.00
    xDefault, yDefault, zDefault = 0.00, 0.00, 1.00
    width, height = 1024, 1024
)

var (
    err error
    x, y, z float64
    bfx, bfy, bfz *big.Float
    xString, yString, zString string
    bfxminDefault = big.NewFloat(xminDefault)
    bfyminDefault = big.NewFloat(yminDefault)
    bfxmaxDefault = big.NewFloat(xmaxDefault)
    bfymaxDefault = big.NewFloat(ymaxDefault)
    bfxDefault = big.NewFloat(xDefault)
    bfyDefault = big.NewFloat(yDefault)
    bfzDefault = big.NewFloat(zDefault)
    bfwidth, bfheight = big.NewFloat(width), big.NewFloat(height)
    xmin, xmax, ymin, ymax = xminDefault, xmaxDefault, yminDefault, ymaxDefault
    bfxmin = big.NewFloat(xmin)
    bfymin = big.NewFloat(ymin)
    bfxmax = big.NewFloat(xmax)
    bfymax = big.NewFloat(ymax)
)

func main() {
    http.HandleFunc("/", mandelHand)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelHand(w http.ResponseWriter, r *http.Request) {
    // ------------------x,y,zoom handling with URL strings---------------
    if err = r.ParseForm(); err != nil {
        log.Print(err)
    }
    xString = r.Form.Get("x")
    if xString != "" {
        if x, err = strconv.ParseFloat(xString, 64); err != nil {
            log.Print(err)
        }
        bfx = big.NewFloat(x)
    } else {
        bfx = big.NewFloat(xDefault)
    }
    yString = r.Form.Get("y")
    if yString != "" {
        if y, err = strconv.ParseFloat(yString, 64); err != nil {
            log.Print(err)
        }
        bfy = big.NewFloat(y)
    } else {
        bfy = big.NewFloat(yDefault)
    }
    zString := r.Form.Get("z")
    if zString != "" {
        if z, err = strconv.ParseFloat(zString, 64); err != nil {
            log.Print(err)
        }
        bfz = big.NewFloat(z)
        // zooming
        bfxmin.Quo(bfxminDefault, bfz)
        bfxmax.Quo(bfxmaxDefault, bfz)
        bfymin.Quo(bfyminDefault, bfz)
        bfymax.Quo(bfymaxDefault, bfz)
    } else {
        bfz = big.NewFloat(zDefault)
    }

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        // OMG: lack of operator overloading
        bfpy := big.NewFloat(float64(py))
        bfyy := big.NewFloat(0)
        omg := big.NewFloat(0)
        omg.Sub(bfymax, bfymin)
        omg.Mul(omg, bfheight)
        bfyy.Quo(bfpy, omg)
        bfyy.Add(bfyy, bfymin)
        bfyy.Add(bfyy, bfy)
        //yy := (float64(py) / height * (ymax - ymin) + ymin) + y
        for px := 0; px < width; px++ {
            bfpx := big.NewFloat(float64(px))
            bfxx := big.NewFloat(0)
            omg1 := big.NewFloat(0)
            omg1.Sub(bfxmax, bfxmin)
            omg1.Mul(omg1, bfwidth)
            bfxx.Quo(bfpx, omg1)
            bfxx.Add(bfxx, bfxmin)
            bfxx.Add(bfxx, bfx)
            //xx := (float64(px) / width * (xmax - xmin) + xmin) + x
            bfzz := MyComplex{Real: bfxx, Imag: bfyy}
            //zz := complex(xx, yy)

            // img point (px, py) represents complex value of z
            img.Set(px, py, mandelbrot(bfzz))
            //img.Set(px, py, mandelbrot(zz))
        }
    }
    png.Encode(w, img)
}

func mandelbrot(z MyComplex) color.Color {
    // those are completely arbitrally values
    const (
        iterations = 0xff
        contrast = 0x00aff0
        alfa = 0xff
    )

    v := NewMyComplex(0, 0)
    black := color.RGBA{0x00, 0x00, 0x00, alfa}
    c := 0x000000
    redmask := 0xff0000
    greenmask := 0x00ff00
    bluemask := 0x0000ff
    var red, green, blue uint8
    for n := uint8(0); n < iterations; n++ {
        c += contrast
        v.Mul(v, v)
        v.Add(v, z)
        var ab *big.Float = v.Abs()
        two := big.NewFloat(float64(2))
        if ab.Cmp(two) > 1 {
            red = uint8((c & redmask) >> 16)
            green = uint8((c & greenmask) >> 8)
            blue = uint8(c & bluemask)
            return color.RGBA{red, green, blue, uint8(alfa)}
        }
    }
    return black
}
