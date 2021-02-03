package main

import (
	"fmt"
)

/*
func f0(output chan int, value int) {
    output <- value
}
*/

func f1(input chan int) {
	fmt.Println(<-input)
}

func main() {
	intchan := make(chan int)
	intchan <- 2 // here writing to (unbuffered) channel BLOCKS main thread!!! -> deadlock
	go f1(intchan) // this line has no chance to be executed, thus consumer does not consume

    /* attempts to solve this deadlock:

    1. put go f1(intchan) BEFORE intchan <- 2 ---> put consumer before producer
        intchan := make(chan int)
        go f1(intchan)                               // this is separate thread
        intchan <- 2

    2. put go f0 into the code ------> consumer and producer are in seperate goroutines
        intchan := make(chan int)
        go func(output chan int, value int) {
            output <- value
        }(intchan, 2)
        go f1(intchan)

    3. make buffered channel instead of unbuffered
        intchan := make(chan int, 2)
        intchan <- 2
        go f1(intchan)
    */
}
