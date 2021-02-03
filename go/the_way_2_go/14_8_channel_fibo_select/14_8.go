/*
Not very efficient (because of transmission complexity) fibonacci program
with extensive goroutines and channels usage.
*/
package main

import (
    "fmt"
    "time"
)

const SECOND time.Duration = 1e9

func main() {
    printsFibo(fiboGenerates(), quits())
}

// sends fibonacci series into return channel(read only)
func fiboGenerates() <-chan int {
    ch := make(chan int)

    go func() {
        first, second := 0, 1
        fibo := first + second
        ch <- first
        ch <- second

        for {
            ch <- fibo
            first, second = second, fibo
            fibo = first + second
        }
    }()

    return ch
}

// send quit signal to channel
func quits() <-chan bool {
    q := make(chan bool)

    go func() {
        time.Sleep(12 * SECOND)                                    // time bomb
        q <- true
    }()

    return q
}

func printsFibo(fibo <-chan int, quit <-chan bool) {
    var v int
    for {
        select {
        case v = <-fibo:
            time.Sleep(SECOND)
            fmt.Println(v)
        case <-quit:
            fmt.Println("It's time to quit.")
            return
        }
    }
}
