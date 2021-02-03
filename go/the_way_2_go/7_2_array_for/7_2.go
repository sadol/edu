// FOR-RANGE idiom used with an array

package main

import "fmt"

func main() {
	const size int = 20
	var arr1 [size]int
    fmt.Println("Newly created array is automatically initialized:")
    for i := range arr1 {
		fmt.Printf("arr[%d] = %d.\n", i, arr1[i])
	}
    fmt.Println()
    fmt.Println("New values in the array:")
	for i := range arr1 {
		arr1[i] = i
		fmt.Printf("arr[%d] = %d.\n", i, arr1[i])
	}
}
