// slices: reference type (more C array type than array itself) => never use
// pointer to a slice

package main

import "fmt"

func main() {
	const maxLen = 50
	fiboArrSlice := make([]int, maxLen)     // declare array + make slice of it + return slice object (not a pointer)
	initFibo(fiboArrSlice)
	fmt.Println(fiboArrSlice)
	maxSlice := 20
	fiboSlice := initFiboPart(maxSlice)
	fmt.Println(fiboSlice)
}

func initFibo(fiboSlice []int) {                         // this is a slice !!!
	for i := range fiboSlice {
        if fiboSlice[i] != 0 { continue }
		if i == 0 || i == 1 {
			fiboSlice[i] = 1
		} else {
			fiboSlice[i] = fiboSlice[i-2] + fiboSlice[i-1]
		}
	}
}

func initFiboPart(maxSlice int) (retSlice []int) {
	if maxSlice < 0 || maxSlice > 50 {
		retSlice = nil
	} else {
		retSlice = make([]int, maxSlice)
		initFibo(retSlice)
	}
	return
}
