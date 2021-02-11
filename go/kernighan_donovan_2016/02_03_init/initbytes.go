package main

import (
    "fmt"
    "./popcount"
    "time"
)

func main() {
    tests := []uint64{1111, 234567, 123456789, 345678910, 45678123456}

    fmt.Println()
    start := time.Now()
    for _, val := range tests {
        popcount.PopCount(val)
    }
    fmt.Printf("PopCount time valuation: %v.\n", time.Since(start))

    start = time.Now()
    for _, val := range tests {
        popcount.PopCountLoop(val)
    }
    fmt.Printf("PopCountLoop time valuation: %v.\n", time.Since(start))
}
