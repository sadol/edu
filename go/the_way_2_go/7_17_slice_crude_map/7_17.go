// creating map function using some lambda

package main

import "fmt"

func main() {
    test := []int{1, 2, 3, 4, 5, 6, 7, 8}
	retSlice := mapFunc(lambda10, test)
	fmt.Println("Original array:", test, ", mapped array:", retSlice, ".")
}

func mapFunc(f func(int) int, input []int) []int {
	returnSlice := make([]int, len(input))
	for index, value := range input {
		returnSlice[index] = f(value)
	}
	return returnSlice
}

func lambda10(input int) int {
	return input * 10
}
