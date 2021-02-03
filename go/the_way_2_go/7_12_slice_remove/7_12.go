// removing subslice from a slice

package main

import "fmt"

func main() {
	inputString := make([]byte, 10, 30)
	inputString[0] = 'a'
	inputString[1] = 'b'
	inputString[2] = 'c'
	inputString[3] = 'd'
	inputString[4] = 'e'
	inputString[5] = 'f'
	inputString[6] = 'g'
	inputString[7] = 'h'
	inputString[8] = 'i'
	inputString[9] = 'j'
	fmt.Printf("Base byte string: %q.\n ", inputString)
    from := 3
    to := 6
	fmt.Printf("Remove bytes from %d to %d.\n ", from, to)
    if byteStringReturn, ok := RemoveStringSlice(inputString, from, to); ok {
		fmt.Printf("Modified byte string: %q.\n ", byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}
}

// Removes subslice froma slice
func RemoveStringSlice (inputSlice []byte, from int, to int) (returnSlice []byte, ok bool) {
    if from < 0 || to < 0 || from > to || to > len(inputSlice) {
        ok = false
    } else {
        returnSlice = make([]byte, len(inputSlice) + from - to)
        copy(returnSlice[:from], inputSlice[:from])
        copy(returnSlice[from:], inputSlice[to:])
        ok = true
    }
    return
}
