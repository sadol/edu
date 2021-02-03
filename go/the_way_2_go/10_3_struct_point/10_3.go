// standard `give me a point' exercise; golang initial version

package main

import (
	"fmt"
    "./point"
)

func main() {
	p1 := point.Point{X: 23.45, Y: 12.4}
	fmt.Println(p1)
	abs := point.Abs(&p1)
	fmt.Println(abs)

	p2 := &point.Point{X: 1, Y: 1}
	fmt.Println(*p2)
	abs2 := point.Abs(p2) // no need to use arraow operator or reference like in C
	fmt.Println(abs2)
	p2 = point.Scale(p2, 2)
	fmt.Println(*p2)
	abs2 = point.Abs(p2)
	fmt.Println(abs2)
}
