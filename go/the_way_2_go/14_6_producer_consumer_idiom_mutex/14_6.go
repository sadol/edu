package main

import (
	"./semaphore"
	"fmt"
	"time"
)

// puts series of ints int channel
func producer(ch chan int) {
    for i := 0; i < 100; i += 10 {
        ch <- i
    }
    close(ch)                                    // give signal to the consumer
}

// removes ints from channel
func consumer(ch chan int, mxComplete semaphore.Semaphore) {
    for i := range ch {  // recieving signal from `close' in THIS particular idiom!!! → like in other (more COMPLEX) idiom: i, close := <- ch + test close...
	    time.Sleep(1e9)
        fmt.Println(i)
    }
    mxComplete.Signal()                  // give signal of end of computations
}

func main () {
    const LEN = 3
    fmt.Println("main() starts.")
    chResult := make(chan int, LEN)
    mutexComplete := semaphore.NewSemaphore()

    go consumer(chResult, mutexComplete)
    go producer(chResult)

    mutexComplete.Wait(1)
    fmt.Println("main() concludes.")
}
