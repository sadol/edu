package main

import (
	"fmt"
	"time"
)

const second time.Duration = 1e9 // number of nanoseconds in one second

// puts integers into the channel `ch'
func producer(ch chan int, steps int) {
	for i := 0; i < steps; i++ {
		fmt.Printf("Putting %v into the channel.\n", i)
		ch <- i // golang idiom for putting into the channel
		//time.Sleep(second) // primitive blocking
	}
}

// pulls integers from the channel `ch'
func consumer(ch chan int) {
	for {
		fmt.Printf("Taking %v from the channel.\n", <-ch) // golang idiom for pulling from the channel, using & discariding
		//time.Sleep(second)                                // primitive blocking
	}
}

func main() {
	fmt.Println("Staring the main")
	steps, buffLen := 10, 5
	blockingChan := make(chan int, buffLen) // channels are referenced types
	go consumer(blockingChan)
	go producer(blockingChan, steps)
	time.Sleep(second) // "server loop" stub
	fmt.Println("Stopping the main")
}
