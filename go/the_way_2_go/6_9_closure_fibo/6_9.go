// closured fibonacci

package main

import "fmt"

func main() {
	f := fibo2()
	for i := 0; i < 10; i++ {
		fmt.Printf("id: %d\t, fibo: %d\n", i+1, f())
	}
}

func fibo2() func() int {
	previous := 0
	current := 1
	return func() int {
		next := previous + current
		ret := current
		previous = current
		current = next
		return ret
	}
}
