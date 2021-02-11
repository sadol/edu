package popcount

import "fmt"

// pc[i] is a population count of `i'
var pc[256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i / 2] + byte(i & 1)
    }
}

func PrintTable() {
    fmt.Println()
    for id, val := range pc {
        if (id % 8) == 0 {
            fmt.Println()
        }
        fmt.Printf("%04b\t", val)
    }
    fmt.Println()
}

func PopCount(x uint64) int {
    return int(
        pc[byte(x >> (0 * 8))] +
        pc[byte(x >> (1 * 8))] +
        pc[byte(x >> (2 * 8))] +
        pc[byte(x >> (3 * 8))] +
        pc[byte(x >> (4 * 8))] +
        pc[byte(x >> (5 * 8))] +
        pc[byte(x >> (6 * 8))] +
        pc[byte(x >> (7 * 8))])
}

func PopCountLoop(x uint64) (ret int) {
    for i := 0; i < 8; i++ {
        ret += int(pc[byte(x >> (i * 8))])
    }
    return
}

func PopCount64shift(x uint64) (ret int) {
    for y := x; y > 0; y >>= 1 {
        if (y & 1) == 1 {
            ret++
        }
    }
    return
}
