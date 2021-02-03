// nested loops

package main

import (
	"fmt"
	// "strings"
)

func main() {
	const loops int = 25
	loop2(loops)
	fmt.Println()
	loop1(loops)
}

// 2 nested loops
func loop2(loops int) {
	for i := 1; i <= loops; i++ {
		fmt.Printf("%d  :", i)
		for j := 0; j < i; j++ {
			fmt.Printf("G")
		}
		fmt.Printf("\n")
	}
}

// only one loop; should use strings.Repeat() instead
func loop1(loops int) {
	G := "G"
	for i := 1; i <= loops; i++ {
		fmt.Printf("%d  : %s\n", i, G)
		G += "G"
	}
}
