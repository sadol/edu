package main

import (
	"./float64array"
    "sort"
	"fmt"
)

func main() {
	const size = 10
    example1 := float64array.NewFloat64Array()
    example2 := float64array.NewFloat64Array(size)
    fmt.Println("Unsorted:")
	fmt.Printf("%v\n", example1)
	fmt.Printf("%v\n", example2)
    fmt.Println()
	sort.Sort(example1) // implicit Sorter interface in action
	sort.Sort(example2)
	fmt.Println("Sorted:")
	fmt.Printf("%v\n", example1)
	fmt.Printf("%v\n", example2)
}
