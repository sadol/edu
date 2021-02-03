// blocking nature of default channels

package main

import (
	"fmt"
	"time"
)

const second time.Duration = 1e9 // number of nanoseconds in one second (1 *10^9)

// puts integers into the channel `ch'
func producer(ch chan int, steps int) {
	for i := 0; i < steps; i++ {
		fmt.Printf("Putting %v into the channel.\n", i)
		ch <- i            // golang idiom for putting into the channel
		time.Sleep(second) // primitive blocking
	}
}

// pulls integers from the channel `ch'
func consumer(ch chan int) {
	for { // golang idiom of emptying the channel, works fine because of blocking nature of recieveing end of the channel
		fmt.Printf("Taking %v from the channel.\n", <-ch) // golang idiom for pulling from the channel, using & discariding
	}
}

func main() {
	fmt.Println("Main starts.")
	blockingChan := make(chan int) // channels are referenced types
	steps := 10
	go consumer(blockingChan)
    //time.Sleep(15 * second)
	go producer(blockingChan, steps)
	time.Sleep(10 * second) // "server loop" stub
	fmt.Println("Main concludes.")
}
