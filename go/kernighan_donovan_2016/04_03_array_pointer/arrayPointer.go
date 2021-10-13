// reversing elements of the array using pointer to an array
package main

import (
    "fmt"
)

// not very versatile use of pointers
func reverse(pArr *[11]int) [11]int {
    for i, j := 0, len(*pArr) - 1; i < j; i, j = i + 1, j - 1 {
        pArr[i], pArr[j] = pArr[j], pArr[i]
    }
    return *pArr
}

func main() {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Printf("\nbefore reverse: %v\n", arr)
    fmt.Printf("after reverse: %v\n\n", reverse(&arr))
}
