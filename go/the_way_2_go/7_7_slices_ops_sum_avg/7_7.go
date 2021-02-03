// summing values in arrays and slices using functions

package main

import "fmt"

func main() {
	var slice = []float64{1.01, 2.01, 3.01, 4.01, 5.01, 6.01, 7.01, 8.01, 9.01, 10.01}
	fmt.Println("Original slice")
	fmt.Println(slice)
	fmt.Printf("is of type %T.\n", slice)
	fmt.Printf("Sum of the all elements (slice version): %f.\n", SumSlice(slice))
	slice1 := slice[:5]
	fmt.Println("Slice")
	fmt.Println(slice1)
	fmt.Printf("is of type %T.\n", slice1)
	fmt.Printf("Sum of the all elements (slice version): %f.\n", SumSlice(slice1))
	// barebone array (rarely used in golang, unlike slices)
	var arr2 [4]float64
	arr2[0] = 1.01
	arr2[1] = 2.01
	arr2[2] = 3.01
	arr2[3] = 4.01
	fmt.Println("4 element array")
	fmt.Println(arr2)
	fmt.Printf("is of type %T.\n", arr2)
	fmt.Printf("Sum of the all elements (array version): %f.\n", SumArr(arr2))
	sum, average := SumAndAverage(slice)
	fmt.Printf("Sum of the all elements %f and average is %f .\n", sum, average)
    fmt.Println()
    fmt.Println("And now it's a time for variadic function usage: ")
	fmt.Printf("Sum of the all elements (variadic version with slice...): %f.\n", SumVariadic(slice...))
    fmt.Println()
	//fmt.Printf("Sum of the all elements (variadic version with array...): %f.\n", SumVariadic(arr2...))
    fmt.Println()
}

// 4 elment array version
func SumArr(array [4]float64) float64 {
	sum := 0.0
	//for i := 1; i < 4; i++ {
    for i := range array { // presence of only one variable means that 'i' is index (not value)
		sum += array[i]
	}
	return sum
}

// slice version, no explicit dimention needed
func SumSlice(array []float64) float64 {
	sum := 0.0
	for _, value := range array { // ignoring indexes, using only values
		sum += value
	}
	return sum
}

// better version -> variadic function
func SumVariadic(elements ...float64) (sum float64) {
    for _, val := range elements { sum += val }
    return
}

func SumAndAverage(slice []float64) (float64, float64) {
    sum, average := SumSlice(slice), .0 // go compiler automatically recognizes .0 as a float
	average = sum / float64(len(slice))
	return sum, average
}
