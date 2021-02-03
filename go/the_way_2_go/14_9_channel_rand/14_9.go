// random bit generator using goroutines and channels
package main

import (
    "fmt"
    "time"
    "math/rand"
)

const (
    LINELEN int = 40                              // default lenght of the line
    QUANT time.Duration = 1e8                           // 1/10th of the second
)

// binary randomizer function
func randomizes01() <-chan int {
    ch := make(chan int)

    go func() {
        for {
            time.Sleep(QUANT)
	        rand.Seed(time.Now().UTC().UnixNano())   // concurrent safe version
            ch <- rand.Intn(2)
        }
    }()

    return ch
}

// printer function for randomizer
func prints(ch <-chan int) {
    for {
        for i := 0; i < LINELEN; i++ {
            fmt.Printf("%d", <-ch)
        }
        fmt.Println()
        time.Sleep(2 * QUANT)
    }
}

func main() {
    prints(randomizes01())
}
