// another look at packages management

package main

import (
	"./even"
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if even.IsEven(i) {
			fmt.Printf("Number %d is even.\n", i)
		} else {
			fmt.Printf("Number %d is odd.\n", i)
		}
	}
}
