// slice rotation in single pass
package main

import "fmt"

func rotate(slice []int, position int) []int {
    var pos int
    switch {
    case (position % len(slice)) == 0:
        return slice
    case position < 0:
        pos = len(slice) + position
    case position > 0:
        pos = position
    }

    ret := make([]int, len(slice))
    copy(ret[:pos], slice[len(slice) - pos:])
    copy(ret[pos:], slice[:len(slice) - pos])
    return ret
}

func main() {
    sl := []int{0,1,2,3,4,5,6,7,8,9}
    pos1, pos2 := 3, -3
    fmt.Printf("\ninitial slice: %v\n", sl)
    fmt.Printf("slice rotated by factor of %d: %v\n", pos1, rotate(sl, pos1))
    fmt.Printf("slice rotated by factor of %d: %v\n", pos2, rotate(sl, pos2))
}
