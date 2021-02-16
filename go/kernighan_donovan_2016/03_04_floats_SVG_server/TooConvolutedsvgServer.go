// computes svg rendering 3D surface function
// USAGE: `localhost:8000/?h=<some height>&w=<some width>&c=<some color code>'
//                       ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//  URL query string, can bypass html alltogether
package main

import (
    "io"
    "fmt"
    "math"
    "sort"
    "net/http"
    "log"
    "strconv"
)

const (
    cells         = 100                                 // number of grid cells
    angle         = math.Pi / 6                   // angle of x , y axes (=30°)
    afactor       = 0.5                  // `a' factor for the egg-box function
    bfactor       = 2.0                  // `b' factor for the egg-box function
)

var (
    sin30 = math.Sin(angle)
    cos30 = math.Cos(angle)
)

func main() {
    http.HandleFunc("/", svgHand)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// main handler
func svgHand (w http.ResponseWriter, r *http.Request) {
    var err error
    var h, we, c, polyColor int
    var h64, we64, c64 int64
    width, height := 600, 320                           // canvas size in pixels
    // extract parameters from `r' here ...
    if err = r.ParseForm(); err != nil {
        log.Print(err)
    }

    hString := r.Form.Get("h")
    if hString == "" {
        h = height
    } else {
        h64, err = strconv.ParseInt(hString, 10, 0)
        h = int(h64)
        if err != nil { log.Print(err) }
    }
    wString := r.Form.Get("w")
    if wString == "" {
        we = width
    } else {
        we64, err = strconv.ParseInt(wString, 10, 0)
        we = int(we64)
        if err != nil { log.Print(err) }
    }
    cString := r.Form.Get("c")
    if cString != "" {
        c64, err = strconv.ParseInt(cString, 16, 0)
        c = int(c64)
        if err != nil { log.Print(err) }
    }

    fmt.Println(h, we, c)  // DEBUG
    // stores isometric coordinates of SVG's polygons
    isoValues := make([]map[string]float64, cells * cells)
    fillValues(we, h, isoValues, f)
    minimum, maximum := findExtrema(isoValues)

    w.Header().Set("Content-Type", "image/svg+xml")        // IMPORTANT for SVG

    // construct SVG's polygon
    io.WriteString(w, fmt.Sprintf("<svg xmlns='http://www.w4.org/2000/svg' " +
                      "style='stroke: grey; fill: none; stroke-width: 0.7' " +
                      "width='%d' height='%d'>", we, h))
    for _, v := range isoValues {
        if math.IsInf(v["az"], 0) || math.IsInf(v["bz"], 0) ||
           math.IsInf(v["cz"], 0) || math.IsInf(v["dz"], 0) {
            continue
        }
        if c != 0 {
            polyColor = c
        } else {
            polyColor = giveColor(minimum, maximum, v["az"], v["bz"], v["cz"], v["dz"])
        }
        io.WriteString(w, fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='#%06x'/>\n",
                       v["ax"], v["ay"], v["bx"], v["by"], v["cx"], v["cy"],
                       v["dx"], v["dy"], polyColor))
    }
    io.WriteString(w, "</svg>")
}

func corner(i, j int, f func(float64, float64) float64, w, h int) (float64, float64, float64) {
    xyrange := 30.0                    // axis ranges (-xyrange..+xyrange)
    xyscale := w / 2 / int(xyrange)               // pixels per x or y unit
    zscale := float64(h) * 0.4              //     pixels per z unit
    // find point (x, y) at the corner of the cell (i, j)
    x := xyrange * (float64(i) / cells - .5)
    y := xyrange * (float64(j) / cells - .5)

    // compute surface hight z
    z := f(x, y)                 // HERE you can put eggbox or saddle functions

    // project (x, y, z) isometrically into 2d svg canvas (sx, sy)
    sx := float64(w) / 2 + (x - y) * cos30 * float64(xyscale)
    sy := float64(h) / 2 + (x + y) * sin30 * float64(xyscale) - z * float64(zscale)
    return sx, sy, z
}

// fills proper slices with isometric coordinates for later use by other
// functions
func fillValues(w, h int, table []map[string]float64, f func(float64, float64) float64) {
    var k int
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            table[k] = make(map[string]float64, 8)
            table[k]["ax"], table[k]["ay"], table[k]["az"] = corner(i + 1, j, f, w, h)
            table[k]["bx"], table[k]["by"], table[k]["bz"] = corner(i, j, f, w, h)
            table[k]["cx"], table[k]["cy"], table[k]["cz"] = corner(i, j + 1, f, w, h)
            table[k]["dx"], table[k]["dy"], table[k]["dz"] = corner(i + 1, j + 1, f, w, h)
            k++
        }
    }
}

// very crude way to find approximate extrema just for graph coloring scale
// setup
func findExtrema(table []map[string]float64) (minimum, maximum float64) {
    for _, v := range table {
        if math.IsInf(v["az"], 0) || math.IsInf(v["bz"], 0) ||
           math.IsInf(v["cz"], 0) || math.IsInf(v["dz"], 0) {
            continue
        }
        sorted := []float64{v["az"], v["bz"], v["cz"], v["dz"]}
        sort.Float64s(sorted)
        if sorted[0] < minimum { minimum = sorted[0] }
        if sorted[3] > maximum { maximum = sorted[3] }
    }
    return
}

// computes proper color to fill a polygon
func giveColor(minimum, maximum, az, bz, cz, dz float64) (color int) {
    maxColor, zeroColor := 0xff0000, 0x00ff00
    // minColor = 0x0000ff
    positiveStep := 0x000100                        // how far to the next color code
    //negativeStep := 0x000001                        // how far to the next color code
    numberOfColors := func () int {          // number of colors in the upper or lower palette
        var combinations int
        for i := zeroColor + positiveStep; i <= maxColor; i += positiveStep {
            combinations++
        }
        return combinations
    }()
    //mean := (az + bz + cz + dz) / 4.0
    maxTable := []float64{az, bz, cz, dz}
    sort.Float64s(maxTable)
    maxLocal := maxTable[3]
    switch {
    case maxLocal > 0.0:
        color = int((maxLocal / maximum) * float64(numberOfColors))
        color = color * positiveStep
    case maxLocal < 0.0:
        color = int((maxLocal / minimum) * float64(numberOfColors))
        // color = int(color) * negativeStep
    case maxLocal == 0.0:
        color = zeroColor
    }
    return
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

// additional functions:
func eggBox(x, y float64) float64 {
    return afactor * (math.Sin(x / bfactor) + math.Sin(y / bfactor))
}

func saddle(x, y float64) float64 {
    return  (x * x * x * x) - 2 * (x * x) + (y * y)
}
