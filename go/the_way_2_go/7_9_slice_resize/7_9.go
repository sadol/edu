// reslicing a slice

package main

import "fmt"

func main() {
	const myLen = 20
	const myCap = 40
	mySlice := make([]int, myLen, myCap)
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3
	mySlice[3] = 4
	mySlice[4] = 5
	mySlice[5] = 6
	mySlice[6] = 7
	mySlice[7] = 8
	fmt.Println("Original slice:")
	fmt.Println(mySlice)
	fmt.Printf("Length of original slice is %d and capacity is %d.\n",
		len(mySlice), cap(mySlice))
	mySlice = magnifySlice(mySlice, 2)
	fmt.Println("Enlarged (twice) slice:")
	fmt.Println(mySlice)
	fmt.Printf("Length of enlarged slice is %d and capacity is %d.\n",
		len(mySlice), cap(mySlice))
    mySlice = magnifySlice(mySlice, 2)  // once again
	fmt.Println("Enlarged (quadrupled) slice:")
	fmt.Println(mySlice)
	fmt.Printf("Length of enlarged slice is %d and capacity is %d.\n",
		len(mySlice), cap(mySlice))
}

// reslicing
func magnifySlice(slice []int, enlargeFactor int) []int {
    returnSlice := make([]int, enlargeFactor*len(slice), enlargeFactor*cap(slice))
	_ = copy(returnSlice, slice) // ignore number of elements copied info
	slice = returnSlice
	return slice
}
