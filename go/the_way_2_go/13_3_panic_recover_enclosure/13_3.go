package main

import (
    "errors"
	"fmt"
	"log"
    "math"
)

// on 64 bit machines int<=>int64 so no panic is expected there
func ConvertInt64ToInt(input int64) int {
    if input >= math.MaxInt32 || input <= math.MinInt32 {
        panic("PANIC")
    }
    return int(input) // in case of panic this function never return properly
}

func IntFromInt64(input int64) (output int, err error) { // panics management
	defer func() {
        if e := recover(); e != nil {
            err = errors.New("SHIT!")
            log.Printf("Panic catched → value out of range: %v", input)
        } else {
            err = nil // do not count on explicit `err' initialization
        }
	}()
	output = ConvertInt64ToInt(input)
	return // even in case of panic recovery this function returs 0 !!!
}

func main() {
	testCases := []int64{12, 100000000, 3000000000000000000, 1}
	for _, caset := range testCases {
        if value, err := IntFromInt64(caset); err == nil {
            // despite panic recovery there is need to proper outcome presentation
            // to the user → output value of 0 is suspicious
	        fmt.Printf("Case: int64(%v) gives: int(%v).\n", caset, value)
        }
	}
    fmt.Println("Main concludes.")
}
