// functions of functions of slices (higher order functions)

package main

import "fmt"

func main() {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice := arr1[2:6]
	filteredSlice := Filter(mySlice, checkOddity)
	fmt.Println("Original base array:")
	fmt.Println(arr1)
	fmt.Printf("Original array len is %d and capacity is %d.\n", len(arr1), cap(arr1))
	fmt.Println()
	fmt.Println("Original slice:")
	fmt.Println(mySlice)
	fmt.Printf("Original slice len is %d and capacity is %d.\n",
		len(mySlice), cap(mySlice))
	fmt.Println()
	fmt.Println("Processed slice:")
	fmt.Println(filteredSlice)
	fmt.Printf("Processedd slice len is %d and capacity is %d.\n",
		len(filteredSlice), cap(filteredSlice))
}

// filter function for integer slice, returns filtered slice
func OldFilter(slice []int, f func(int) bool) []int {
	returnSlice := make([]int, 0, cap(slice)) // cap here is necessary for reslicing
	for _, value := range slice {
		if f(value) {
			returnSlice = returnSlice[:len(returnSlice)+1] //reslice!!! --> not very efficient
			returnSlice[len(returnSlice)-1] = value
		}
	}
	return returnSlice
}

// much simpler version
func Filter(slice []int, f func(int) bool) (returnSlice []int) {
	returnSlice = make([]int, len(slice))
    i := 0
	for _, value := range slice {
		if f(value) {
			returnSlice[i] = value
            i++
		}
	}
    return returnSlice[:i]
}

//returns `true' if `value' is odd, `false' otherwise
func checkOddity(value int) (odd bool) {
	if value%2 != 0 {
		odd = true
	} else {
	    odd = false
    }
    return // only 1 point of return makes debugging easier
}
