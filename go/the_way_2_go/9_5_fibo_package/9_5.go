// testing fibo function

package main

import (
	"./fibo"
	"fmt"
)

var value int64 = 0
var index int = 0

func main() {
	sumFibo := fibo.Fibo("+")
	multiFibo := fibo.Fibo("*")

	for i := 0; i < 100; i++ {
		value, index = sumFibo()
		if value == 0 {
			break
		}
		fmt.Printf("Fibo (sum) index: %d\tfibo (sum) value: %d.\n", index, value)
	}

	fmt.Println()

	for i := 0; i < 100; i++ {
		value, index = multiFibo()
		if value == 0 {
			break
		}
		fmt.Printf("Fibo (multi) index: %d\tfibo (multi) value: %d.\n", index, value)
	}
}
