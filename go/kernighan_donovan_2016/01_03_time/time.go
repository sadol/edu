//exercise 1.3: time difference measurement
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// slower piece of code:
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	stop := time.Now()
	diff := stop.Sub(start)
	fmt.Printf("--->Slower code execution time: %v.", diff)
	fmt.Println()
	// faster piece of code:
	start = time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	stop = time.Now()
	diff = stop.Sub(start)
	fmt.Printf("--->Faster code execution time: %v.", diff)
	fmt.Println()
}
