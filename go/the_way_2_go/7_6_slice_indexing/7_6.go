// splitting buffer into 2 slices

package main

import "fmt"

func main() {
	const bufLen int = 10
	buffer := make([]byte, bufLen) // cap == len
	for i := range buffer {
		buffer[i] = byte(i)
	}
	fmt.Println(buffer)
	const header int = 4                     // prefix len
	prefix, suffix := buffer[:4], buffer[4:] // oneliner
	fmt.Println(prefix)
	fmt.Println(suffix)
}
