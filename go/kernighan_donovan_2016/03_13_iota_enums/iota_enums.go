// enums in golang
package main

import (
    "fmt"
)

const (
    KB float64 = 1000
    MB float64 = KB*KB
    GB float64 = MB*KB
    TB float64 = GB*KB
    PB float64 = TB*KB
    EB float64 = PB*KB
    ZB float64 = EB*KB
    YB float64 = ZB*KB
)

func main() {
    fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
