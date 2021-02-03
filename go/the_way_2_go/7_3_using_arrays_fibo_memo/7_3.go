// using arrays

package main

import (
    "fmt"
    "time"
)

func main() {
	const maxLen = 50
	var fiboArr [maxLen]int64
    // note & at the beginnig of print below
    start := time.Now()
    calcAndPrintFibo(&fiboArr)
    end := time.Now()
    delta := end.Sub(start)
    fmt.Println("Slow fibo delta: %d.", delta)
    fmt.Println()
    fmt.Println("Check if values are preserved after exit from function:")
    fmt.Println(fiboArr)
    var fiboArr2 [maxLen]int64
    start = time.Now()
    calcAndPrintFibo2(&fiboArr2)
    end = time.Now()
    delta = end.Sub(start)
    fmt.Println()
    fmt.Println("Fast fibo delta: %d.", delta)
    fmt.Println()

}

// calculates and prints first 50, stores all numbers in the array sent to
// function by reference, this implies no need to return any values (except
// error codes?)
func calcAndPrintFibo2 (fiboA *[50]int64) {
    for i := range fiboA {
        if fiboA[i] != 0 { continue }
		if i == 1 || i == 0 {
			fiboA[i] = 1
		} else {
			fiboA[i] = fiboA[i-2] + fiboA[i-1]
		}
	}
	fmt.Println(fiboA)
}

func calcAndPrintFibo (fiboA *[50]int64) {
    for i := range fiboA {
        //if fiboA[i] != 0 { continue }
		if i == 1 || i == 0 {
			fiboA[i] = 1
		} else {
			fiboA[i] = fiboA[i-2] + fiboA[i-1]
		}
	}
	fmt.Println(fiboA)
}
