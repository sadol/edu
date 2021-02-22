// PNG image of the Mandelbrot fractal
// INFO: creating ZOOM functionality for Mandelbrot fractal is very resouce
// consumig process (for my server at least). Naive solution for zooming
// is to recalculate each time client sends a request BUT it is resource
// depleting activity. More resonable solution is to store high resolution
// picture of the Mandelbrot fractal somewhere on the server and load pieces of
// it on demand. The best solution in my opinion would be using dedicated
// library of some kind for resizing PNG files.

package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "net/http"
    "log"
    "strconv"
    "io"
)

const static = "./static"

func main() {
    http.HandleFunc("/", mandelHand)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelHand(w http.ResponseWriter, r *http.Request) {
    const (
        zeds = make(map[string]int){"1": 1, "2": 2, "4": 4} // available zoom levels
        xmin, ymin, xmax, ymax, = -2, -2, 2, 2
        xDefault, yDefault, zDefault = -1.00, -1.00, 1
        width, height = 16 * 1024, 16 * 1024            // this is BIG file
        mandelFileName = static + "/w.png"              // do not create BIG file unnecassary
        zoomedHeight, zoomedWidth = 1024, 1024          // only this big part of the fractal will be send to the client
    )

    var (
        err error
        z int
        x, y float64
        xString, yString, zString string
        mandelFile *os.File
    )

    // ------------------x,y,zoom handling with URL strings---------------
    if err = r.ParseForm(); err != nil {
        log.Print(err)
    }
    xString = r.Form.Get("x")
    if xString != "" {
        if x, err = strconv.ParseFloat(xString, 64); err != nil {
            log.Print(err)
        }
        if x > xmax || x < xmin { x = xDefault }                    // silently
    } else {
        x = xDefault
    }
    yString = r.Form.Get("y")
    if yString != "" {
        if y, err = strconv.ParseFloat(yString, 64); err != nil {
            log.Print(err)
        }
        if y > ymax || y < ymin { y = yDefault }                    // silently
    } else {
        y = yDefault
    }
    zString := r.Form.Get("z")
    if zString != "" {
        if z, ok := zeds[zString]; !ok {
            z = zDefault
        }
    }

    if _, err = os.Stat(mandelFileName); err != nil {
        if os.IsNotExist(err) {
            // create a file
            if mandelFile, err = os.Create(mandelFileName); err != nil {
                log.Print(err)
            }
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
            png.Encode(mandelFile, img)
            mandelFile.Close()
        } else {
            log.Print(err)
        }
    }

    mandelFile, err = os.Open(mandelFileName)
    if err != nil {
        log.Fatal(err)
    }
    defer mandelFile.Close()
    if err = zoom(w, mandelFile, z, x, y); err != nil {
        log.Fatal(err)
    }
    // INFO: instead of `log.Fatal...' some http errors should be send to the
    // client
}

func zoom(w http.ResponeWriter, f *os.File, z int, x, y float64) error {
    if fullImg := png.Decode(f); err != nil {
        return err
    }

    var subImg image.Image            // zoomed portion of the oryginal fractal
    py := int((y - ymin) * (height * (ymax - ymin)))
    px := int((x - xmin) * (width * (xmax - xmin)))
    zoomedRect := image.Rectangle{}
    var centerPoint image.Point                      // centerpoint of the zoom

    // in the perfect world every zoom should have the same dimentions and
    // supersampling should have been used as necessary
    switch z {
        case 1:
            subImg = fullImg                        // `x' and `y' don't matter
        case 2:
             
            subImg = fullImg.SubImage()
        case 4:

            subImg = fullImg.SubImage()
    }

    io.Copy(w, subImg)
}

// supersampling for zooming
func supersample(output io.Writer, input io.Reader, ) {
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
