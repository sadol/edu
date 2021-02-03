// computing mins and maxs of values in int slices

package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("Array to experiment onto:")
	fmt.Println(slice)
	fmt.Printf("Minimum : %d, maximum : %d.\n", minSlice(slice), maxSlice(slice))
}

func minSlice(slice []int) int {
	min := 0
	for _, value := range slice {
		if min > value {
			min = value
		}
	}
	return min
}

func maxSlice(slice []int) int {
	max := 0
	for _, value := range slice {
		if max < value {
			max = value
		}
	}
	return max
}
