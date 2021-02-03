// make -> allocates space on the heap for SLICE,MAP OR CHAN ,
// initializes it & returns VALUE, in case of slices it creates SLICE HEADER!!
// new -> allocates space on the heap, returns POINTER to zeroed object

package main

import (
	//"bytes"
	"fmt"
)

func main() {

	const preLen = 3
    const preCap = 4
	firstSlice := make([]byte, preLen, preCap) // byte string to be appended to
	// fill some values but do not fill full len of the slice
	firstSlice[0] = 'a'
	firstSlice[1] = 'l'
	fmt.Println(firstSlice)

	const dataLen = 10
	appendedByteArray := make([]byte, dataLen) // byte string to append; cap == len
	appendedByteArray[0] = 'o'
	appendedByteArray[1] = 'o'
	fmt.Println(appendedByteArray)

	// both slices can be wider than number of filled bytes in them !
	if retSlice, valid := Append(firstSlice, appendedByteArray); valid {
		fmt.Println(retSlice)
    } else {
		fmt.Println("Some error ocurred;<")
    }

    fmt.Println()
    fmt.Println("Sparse appending:")
    sparseSlice := AppendSparse(firstSlice, appendedByteArray)
    fmt.Println(sparseSlice)
}

// appending only non default values (let's assume that 0 is not allowed)
func Append(slice, data []byte) (combinedArray []byte, valid bool) {
	// check amount of free space in the slice:
	numOfCharsInSlice, numOfCharsInData := 0, 0
	for i := range slice {
		if slice[i] != 0 {
			numOfCharsInSlice++
		}
	}
	for i := range data {
		if data[i] != 0 {
			numOfCharsInData++
		}
	}

	numFreeCharsInSlice := cap(slice) - numOfCharsInSlice // amount of free room in a slice
	if numFreeCharsInSlice < numOfCharsInData {        // not enough room in slice buffer
		// create new slice and copy accordingly from both buffers
		combinedArray = make([]byte, numOfCharsInSlice+numOfCharsInData) // cap == len
	} else {
		combinedArray = make([]byte, cap(slice)) // cap == len
	}
	_ = Copy(combinedArray, 0, slice, 0, numOfCharsInSlice)
	valid = Copy(combinedArray, numOfCharsInSlice, data, 0, numOfCharsInData)
	return
}

// copying only non default values
func Copy(dest []byte, destStart int, source []byte, sourceStart int,
	sourceEnd int) bool {
    // check if there is need to create longer array, if so signal error
	if (cap(dest) - destStart) < (sourceEnd - sourceStart) {
		return false
	}
    // no need to create wider buffer, function may proceed
	for i := sourceStart; i < sourceEnd; i++ { // copy byte by byte
		dest[destStart+i] = source[i]
	}
	return true // success
}

/* Alternalive version of the function `Append'; no apriori assumtions needed.
 * Copies ALL elements of the source slice and appends it at the end of the
 * destination slice making sparse table essentialy.
*/
func AppendSparse(slice, data []byte) (combinedArray []byte) {
    cSlice := cap(slice)
    cData := cap(data)
    lSlice := len(slice)
    lData := len(data)
    combinedArray = make([]byte, cSlice + cData)
    for i := 0; i < lSlice; i++ { combinedArray[i] = slice[i] }
    for i := 0; i < lData; i++ { combinedArray[i + cSlice] = data[i] }
    return
}
