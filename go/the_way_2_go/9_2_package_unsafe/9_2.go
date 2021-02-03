// testing sizes of types

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myInt int = 23
	fmt.Printf("Size of `int' on this machine is %d bytes.\n", unsafe.Sizeof(myInt))
	var myUint uint64 = 23
	fmt.Printf("Size of `uint' on this machine is %d bytes.\n", unsafe.Sizeof(myUint))
	var myFloat64 float64 = 23.0
	fmt.Printf("Size of `float64' on this machine is %d bytes.\n", unsafe.Sizeof(myFloat64))
}
