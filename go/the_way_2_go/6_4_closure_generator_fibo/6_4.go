// fibonacci with closure

package main

import "fmt"

func main() {
	reps := 20
	fun := fiboGenerator()
	fmt.Printf("fiboGenerator test for %d repetitions:\n", reps)
	for pos, val, i := 0, 0, 0; i < reps; i++ {
		pos, val = fun()
		fmt.Printf("Position: %d\t\tvalue: %d\n", pos, val)
	}
}

func fiboGenerator() func() (int, int) {
	currentPosition := 1
	previous := 0
	preprevious := 0
	return func() (position int, value int) {
		if currentPosition <= 1 {
			previous = 1
			value = 1
		} else {
			value = preprevious + previous
			preprevious = previous
			previous = value
		}
		position = currentPosition
		currentPosition++
		return
	}
}
