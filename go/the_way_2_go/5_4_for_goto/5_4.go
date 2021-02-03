// initial fun with for construct
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("Consecutive (for loop) id.:%d\n", i)
	}

	fmt.Println()
	i := 0 // sentinel variable
label:
	fmt.Printf("Consecutive (goto code) id.:%d\n", i)
	i++
	if i > 14 {
		os.Exit(0)
	}
	goto label                                                // this is stupid
}
