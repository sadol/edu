// PI approximation algorithm using channels & goroutines (not very efficient
// version due to ease of calculation & presentation). User aided version.
package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
    "math"
)

// struct to store approx PI values
type Approx struct {
    step int
    value float64
}

const (
    msgHello = "\nWelcome to PI approximation programm!\n"
    msgContinue = "Continue ? <y|n> : "
    msgInfo = "PI approximation of the order of [%d] is : [%f].\n"
    msgBye = "Bye.\n"
)

func main() {
    var (
        tempToken string                             // user's answer container
        chApprox chan Approx
        tempApprox Approx
    )

    fmt.Println(msgHello)
    stdinReader := bufio.NewReader(os.Stdin)              // 1. define a stream
    piScanner := bufio.NewScanner(stdinReader)     // 2. define a bufio.Scanner
    /* ScanLines splitter is implicitly set for scanner to use → there is no
    * need to set this explicitly:
    * piScanner.Split(bufio.ScanLines)  // 3. set splitter function for scanner
    */
    chApprox = approximatePi()  // not very clean solution : no explicite way to stop this subroutine, it ends automatically with the main thread.

    for {
        _, _ = stdinReader.Discard(stdinReader.Buffered())    // clean a stream
        fmt.Printf("%s", msgContinue)

        piScanner.Scan()                                   // 4. using splitter
        tempToken = piScanner.Text()
        if strings.EqualFold(tempToken, "n") {
            fmt.Println(msgBye)
            break
        } else {                                                   // means yes
            tempApprox = <-chApprox
            fmt.Printf(msgInfo, tempApprox.step, tempApprox.value)
            continue
        }
    }
}

// approximation function for PI calculating using channels and goroutines
func approximatePi() (chOutput chan Approx) {
    chOutput = make(chan Approx)

    go func() {
        var lastApproxValue float64
        tempApprox := new(Approx)

        for step := 0; ; step++ {
            lastApproxValue += float64(4) * ( math.Pow(float64(-1), float64(step)) / float64(((2 * step) + 1)))
            tempApprox.step = step
            tempApprox.value = lastApproxValue
            chOutput <- *tempApprox                                // block here
        }
    }()

    return
}
