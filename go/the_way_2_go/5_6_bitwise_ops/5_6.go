// fun with bitwise complementation

package main

import "fmt"

func main() {
	var i, negi, mask uint8
    mask = 15                                                       // 00001111
    for i = 0; i < 10; i++ {
        negi = (^i) & mask                   // remove unnecessary leading bits
		fmt.Printf("Bitwise complement of %04b is %04b.\n", i, negi)
	}
}
