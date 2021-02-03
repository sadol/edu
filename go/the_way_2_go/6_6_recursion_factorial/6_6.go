// standard factorial recursive stuff

package main

import "fmt"

func main() {
	max := 30
	fmt.Printf("Factorial test of the first %d integers:\n", max)
	for i := 0; i < max; i++ {
		fmt.Printf("factorial of %d is %d,\n", i, factorial(i))
	}
	fmt.Printf("Factorial test of the first %d integers(named return version):\n", max)
	for i := 0; i < max; i++ {
		fmt.Printf("factorial of %d is %d,\n", i, factorialNamed(i))
	}
}

func factorial(factor int) int {
	if factor > 12 || factor < 1 {
		return -1 // overflow error for 32 bit ints in go
	} else if factor == 1 || factor == 0 {
		return 1
	} else {
		return factor * factorial(factor-1)
	}
}

// much cleaner solution with only one exit point from the function
func factorialNamed(factor int) (result int) {
	if factor > 12 || factor < 0  {
		result = -1 // overflow error for 32 bit ints in go
	} else if factor == 1 || factor == 0 {
		result = 1
	} else {
		result = factor * factorial(factor-1)
	}
	return
}
