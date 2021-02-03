// second attempt to synch gouroutiens with `semaphore' package (mutexes + chan)
package main

import (
	"./semaphore"
	"fmt"
	"time"
)

// `computational' goroutine function
func sumator(chResultPassing chan int, first, second int) {
	time.Sleep(5 * 1e9)                 // mocking up calculations of some sort
    chResultPassing <- first + second   // critical section + signal completion
}

// printer goroutine function
func printer(chResultPassing chan int, mutexCompleteSignal semaphore.Semaphore) {
    fmt.Printf("Result is %d.\n", <-chResultPassing) // critical section + wait
    mutexCompleteSignal.Signal()                       // inform the main thread
}

func main() {
    fmt.Println("main() starts.")
    chResult := make(chan int)
    mutexComplete := semaphore.NewSemaphore()
	operand1, operand2 := 1, 2

	go sumator(chResult, operand1, operand2)
    go printer(chResult, mutexComplete)

    mutexComplete.Wait(1)                             // waiting fo signal of completion
    fmt.Println("main() concludes.")
}
