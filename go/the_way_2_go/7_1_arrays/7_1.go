// intro to arrays (fixed len, homogenus, value -> not reference type => can be
// defined with NEW)

package main

import "fmt"

func main() {
	const size int = 10
	var arr1 [size]int // name; size; type convention (size must be non negative integer expr)
	arr2 := arr1                                           // THIS IS DEEP COPY !!! in slices this is not the case !!!
	fmt.Printf("arr1 address is %p, arr2 address is %p.\n", &arr1, &arr2)

    arr3 := new([size]int)                             // `new' returns pointer
    fmt.Printf("arr3 object < new([10]int > is if of type of : %T.\n", arr3)
}
