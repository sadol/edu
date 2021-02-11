package main

import (
    "fmt"
    "./popcount"
    "time"
)

func main() {
    tests := []uint64{1111, 234567, 123456789, 345678910, 45678123456}
    results := make([]int, len(tests))

    fmt.Println()
    start := time.Now()
    for id, val := range tests {
        results[id] = popcount.PopCount(val)
    }
    fmt.Printf("PopCount time valuation: %v.\n", time.Since(start))
    fmt.Printf("PopCount results: %v\n", results)

    results = make([]int, len(tests))
    start = time.Now()
    for id, val := range tests {
        results[id] = popcount.PopCountLoop(val)
    }
    fmt.Printf("PopCountLoop time valuation: %v.\n", time.Since(start))
    fmt.Printf("PopCountLoop results: %v\n", results)

    results = make([]int, len(tests))
    start = time.Now()
    for id, val := range tests {
        results[id] = popcount.PopCount64shift(val)
    }
    fmt.Printf("PopCount64shift time valuation: %v.\n", time.Since(start))
    fmt.Printf("PopCount64shift results: %v\n", results)

    results = make([]int, len(tests))
    start = time.Now()
    for id, val := range tests {
        results[id] = popcount.PopCountRightmost(val)
    }
    fmt.Printf("PopCountRightmost time valuation: %v.\n", time.Since(start))
    fmt.Printf("PopCountRightmost results: %v\n", results)
}
