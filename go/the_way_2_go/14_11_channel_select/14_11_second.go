// PI approximation algorithm using channels & goroutines (not very efficient
// version due to ease of calculation & presentation). Multi-goroutine version
// (number of goroutines is the same as number of processors in the machine).
// Very ineffective example of naive optimization.
package main

import (
    "fmt"
    "math"
    "time"
    "runtime"
)

const NUMFACTORS = 20 // number of elements for return from subgoroutine(magic number)

func main() {
    NUMGOROUTINES := runtime.NumCPU()          // number of goroutines employed
    chPiTerms := make(chan float64)                 // partial results reciever
    chCompleted := make(chan bool, NUMGOROUTINES)              // waiting queue
    chTerminate := make(chan bool)         // end of calculations signal buffer
    chFinalPi := make(chan float64)                    // final result reciever

    _ = runtime.GOMAXPROCS(NUMGOROUTINES)
    start := time.Now()
    // start computing in concurrent threads
    for i := 0; i < NUMFACTORS * NUMGOROUTINES; i += NUMFACTORS {
        go calculateTerms(i, i + NUMFACTORS, chPiTerms, chCompleted)
    }

    go approximatePi(chPiTerms, chFinalPi, chTerminate)               // blocks

    // wait for results in the main thread
    for i := 0; i < NUMGOROUTINES; i++ {
        <-chCompleted                                                 // blocks
    }
    // signal completion → terminate approximaxion thread
    chTerminate <- true                                // unblock approximatePi
    stop := time.Now()

    // present results
    fmt.Printf("Final PI approximation for %d threads (lasted %v) is %f.\n",
               NUMGOROUTINES, stop.Sub(start), <-chFinalPi)
}

/*
* Approximator which cumulates partial results from set of producers into one
* final number; due to existence of bottleneck and `race conditions' this
* algorithm is not very efficient and returns slightly different results every
* time.
*/
func approximatePi(chInput <-chan float64, chOutput chan<- float64, chCompleted <-chan bool) {

    var tempPi float64

    for {
        select {
        case term := <-chInput:
            tempPi += term
        case <-chCompleted:
            chOutput <- tempPi
            return
        }
    }
}

/*
* Goroutine function to use in bulk; calculates one type of PI aproximation
* term & puts it into the output channel.
* Parameters:   start → starting index of the PI series
*               stop → stopping index of the PI series
*               chOutput → output channel for consecutive factors
*               chCompleted → completion signal channel
*/
func calculateTerms(start int, stop int, chOutput chan<- float64, chCompleted chan<- bool) {
    for i := start; i < stop; i++ {
         chOutput <- float64(4) * ( math.Pow(float64(-1), float64(i)) / float64(((2 * i) + 1)))
    }
    chCompleted <- true // done (channel is bufferd accordingly so routine terminate immediately
}
