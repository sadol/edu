//

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fv := func() { j := i; j++ }
		fv()
		gv := func(i int) { j := i; j++ }
		gv(i)
		fmt.Printf("%d - fv is of type %T and has value %v;\tgv is of type %T and has value %v.\n",
			i, fv, fv, gv, gv)
	}
}
