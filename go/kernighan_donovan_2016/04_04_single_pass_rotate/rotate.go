// slice rotation in single pass
package main

import "fmt"

func rotate(slice []int, position int) []int {
    slen := len(slice)
    pos := position % slen
    switch {
    case pos == 0:
        return slice
    case pos < 0:
        pos += slen
    }

    ret := make([]int, slen)
    copy(ret[:pos], slice[slen - pos:])
    copy(ret[pos:], slice[:slen - pos])
    return ret
}

func main() {
    sl := []int{0 ,1 , 2, 3, 4, 5, 6, 7, 8, 9}
    pos1, pos2, pos3, pos4, pos5, pos6 := 3, -3, 40, -40, -45, 45
    fmt.Printf("\ninitial slice: %v\n", sl)
    fmt.Printf("slice rotated by factor of %d: %v\n", pos1, rotate(sl, pos1))
    fmt.Printf("slice rotated by factor of %d: %v\n", pos2, rotate(sl, pos2))
    fmt.Printf("slice rotated by factor of %d: %v\n", pos3, rotate(sl, pos3))
    fmt.Printf("slice rotated by factor of %d: %v\n", pos4, rotate(sl, pos4))
    fmt.Printf("slice rotated by factor of %d: %v\n", pos5, rotate(sl, pos5))
    fmt.Printf("slice rotated by factor of %d: %v\n", pos6, rotate(sl, pos6))
}
