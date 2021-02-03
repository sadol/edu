//custom sqrt functions: named & unnamed return vals versions

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	input1, input2 := -12.45, 345.345
	output, err := MySqrt_u(input1)
	fmt.Printf("MySqrt_u(%f) => output=%f\terr=%s\n", input1, output, err)
	output, err = MySqrt_u(input2)
	fmt.Printf("MySqrt_u(%f) => output=%f\terr=%s\n", input2, output, err)
	output, err = MySqrt(input1)
	fmt.Printf("MySqrt(%f) => output=%f\terr=%s\n", input1, output, err)
	output, err = MySqrt(input2)
	fmt.Printf("MySqrt(%f) => output=%f\terr=%s\n", input2, output, err)
}

// unnamed return vals
func MySqrt_u(input float64) (float64, error) {
	if input < 0 {
		return math.NaN(), errors.New("Negative argument forbidden.")
	}
	return math.Sqrt(input), nil
}

// named return vals
func MySqrt(input float64) (output float64, err error) {
	output = math.Sqrt(input)
	if math.IsNaN(output) {
		err = errors.New("Negative argument forbidden.")
	}
	return
}
