// 14.7 c, using second channel for signaling end of the counting
package main

import (
    "fmt"
    "time"
)

const SECOND time.Duration = 1e9

func main() {
    driver(numbers(), quits())
}

// number producer; returns READ ONLY channel (where `chan<- int' is WRITE ONLY idiom)
func numbers() <-chan int {
    ch := make(chan int)

    go func() {
        for i := 0; ;i++ {
            ch <- i
        }
    }()

    return ch
}

// send quit signal to the channel
func quits() <-chan bool {
    quit := make(chan bool)

    go func () {
        time.Sleep(12 * SECOND)
        quit <- true
    }()

    return quit                      // do not wait 12 secs, return imidiatelly
}

// printer function, this is consumer; no need to use biderictional buffer
func driver(ch <-chan int, quit <-chan bool) {
    for {
        select {                                   // `channel switching' idiom
        case v := <-ch:                                       // ususal counter
            time.Sleep(SECOND)
            fmt.Println(v)
        case <-quit:                    // end of counting signal reciever
            fmt.Println("It's high time to quit.")
            return
        }
    }
}
