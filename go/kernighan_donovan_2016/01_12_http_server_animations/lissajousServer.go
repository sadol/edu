// Simple http server to combine gif creation knowledge with some basic
// networking.
// USAGE: `localhost:8000/?cycles=<some value>' into your browswer of choice
//                       ^^^^^^^^^^^^^^^^^^^^^
//  URL query string, can bypass html alltogether
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    //"os"
    "net/http"
    "log"
    "strconv"
)

var (
    black = color.Black
    white = color.White
    green = color.RGBA{0x00, 0xff, 0x00, 0xff}
    red = color.RGBA{0xff, 0x00, 0x00, 0xff}
    blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
)

var palette = []color.Color{black, white, green, red, blue}

const (
    NCOLORS = 4                                   // range of the color palette
)

func main() {
    http.HandleFunc("/", lissaHand)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// main function hadler
func lissaHand(w http.ResponseWriter, r *http.Request) {
    // extract parameters from `r' here ...
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    cycles, err := strconv.ParseFloat(r.Form.Get("cycles"), 64)
    // other parameters of `lissajous' may be serviced the same way with query
    // strings & `strconv'
    if err != nil { log.Print(err) }
    lissajous(w, cycles, 0.001, 100, 64, 8)
}

/*
nframes - number of animation frames  ---> 64
size - image canvas covers [-size..+size]  ---> 100
delay - delay between frames in 10ms units  ---> 8
cycles - number of complete `x' oscillator revolutions ---> 5
res - angular resolution ---> 0.001
*/
func lissajous(out io.Writer, cycles, res float64, size, nframes, delay int) {
    freq := rand.Float64() * 3.0         // relative frequency of `y' oscilator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0                                            // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size +1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles * 2 * math.Pi; t+= res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            img.SetColorIndex(size + int(x * float64(size) + 0.5),
                              size + int(y * float64(size) + 0.5),
                              uint8((i % NCOLORS) + 1))
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
