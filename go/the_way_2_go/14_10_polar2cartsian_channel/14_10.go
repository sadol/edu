// Intreactive CLI program with double channeled goroutine for converting polar
// to cartesian coords (not very efficient way of doing easy calculations using
// dedicated goroutines); program uses bufio.Scan instead of dedicated string
// or byte slice functions because it is safer to use bufio in case of data
// streams processing.
package main

import (
    "fmt"
    "./polar"
    "./cartesian"
    "os"
    "bufio"
    "strconv"
    "strings"
)

const (
    helloMsg = "Welcome to the POLAR TO CARTESIAN CONVERTER!\n"
    coordsMsg = "Please enter polar coordinates(<r> <phi>) or [q] to exit : "
    errorMsg = "Invalid parameter : < %v >.\n"
    byeMsg = "Bye!\n"
)

func main() {
    stdinReader := bufio.NewReader(os.Stdin)              // 1. define a stream
    coordScanner := bufio.NewScanner(stdinReader)  // 2. define a bufio.Scanner
    coordScanner.Split(bufio.ScanWords) // 3. set splitter function for a scanner

    var (
        r float64                                            // `r' polar param
        phi float64                                        // `phi' polar param
        err error                                       // processing error var
        inputPolar polar.Polar                               // temporary value
        tempToken string
    )

    // communication & synchronization artefacts
    input := make(chan polar.Polar)

    // run calculation goroutine here, output channel ready to use
    output := convertToPolar(input)

    fmt.Println(helloMsg)

    for {
        _, _ = stdinReader.Discard(stdinReader.Buffered())    // clean a stream
        fmt.Printf("%s", coordsMsg)

        coordScanner.Scan()                     // 4. using splitter (1st time)
        tempToken = coordScanner.Text()
        if strings.EqualFold(tempToken, "q") {
            fmt.Println(byeMsg)
            break
        } else {
            if r, err = strconv.ParseFloat(tempToken, 64); err != nil {
                fmt.Printf(errorMsg, tempToken)
                continue
            }
        }

        coordScanner.Scan()                     // 4. using splitter (2nd time)
        tempToken = coordScanner.Text()
        if phi, err = strconv.ParseFloat(tempToken, 64); err != nil {
            fmt.Printf(errorMsg, tempToken)
            continue
        }

        // utilize obtained values in some way (not very effective use of
        // goroutine)
        inputPolar = *(polar.NewPolar(r, phi))
        input <- inputPolar
        fmt.Printf("Old polar: %v --->magic---> New cartesian: %v\n\n", inputPolar, <-output) // waits here
    }
}

// goroutine function to converting polar to cartesian coordiantes of the 2D
// point
func convertToPolar(input chan polar.Polar) (output chan cartesian.Cartesian) {
    output = make(chan cartesian.Cartesian)
    var tempPolar polar.Polar    // try not to redefine inside of the goroutine

    go func() {
        for {
            tempPolar = <-input     // wait for input, do not block the main thread
            output <- *tempPolar.PolarToCartesian()
        }
    }()

    return output
}
