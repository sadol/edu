// simple program to convert tempratures (C, K, F)
// in order to feed negative numbers to this program use `--' separator, for
// example `./temps -- -212'
package main

import (
    "flag"
    "fmt"
    "./tempconv"
    "os"
    "strconv"
    "bufio"
)

const (
    TEMPRATURE = "t"
    LENGTH = "l"
    WEIGTH = "w"
    typeERR = "Unknown type <%s>. Legal values of `type': t,l,w.\n"
)

const (
    CELSIUS = "C"
    FAHRENHEIT = "F"
    KELVIN = "K"
    tempUnitERR = "Unknown value <%s> of type `t'. Legal values of `t': C,F,K.\n"
)

// flag buffers
var convType string
var fromVal string
var toVal string

func init () {
    // defining short AND long options alike in `init' because of undefined
    // order of initialization
    const (
        defaultType = TEMPRATURE
        typeUsage = "Conversion type [t (temprature)|l (length)|w (weigth] (temprature by default)"
        defaultFromVal = FAHRENHEIT
        fromValUsage = "Conversion unit of the source (°F by default)"
        defaultToVal = CELSIUS
        toValUsage = "Conversion unit of the destination (°C by default)"
    )
    flag.StringVar(&convType, "type", defaultType, typeUsage)                   // long version
    flag.StringVar(&convType, "T", defaultType, typeUsage + " (shorthand).")    // short version ...
    flag.StringVar(&fromVal, "from", defaultFromVal, fromValUsage)
    flag.StringVar(&fromVal, "f", defaultFromVal, fromValUsage + "(shorthand).")
    flag.StringVar(&toVal, "to", defaultToVal, toValUsage)
    flag.StringVar(&toVal, "t", defaultToVal, toValUsage + " (shorthand).")
}

func main() {
    var rawArgs []string
    var convValues []float64
    var outputValues []float64
    var noArgs int

    flag.Parse()
    if flag.NArg() > 0 {
        noArgs = flag.NArg()
        rawArgs = flag.Args()
    } else {                                          // try to read from stdin
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Split(bufio.ScanWords)
        for scanner.Scan() {
            noArgs++
            rawArgs = append(rawArgs, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    convValues = make([]float64, noArgs)
    outputValues = make([]float64, noArgs)

    for id, val := range rawArgs {
        if val, err := strconv.ParseFloat(val, 64); err == nil {
            convValues[id] = val
        } else {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    // check type
    switch convType {
    case TEMPRATURE:
        // check from
        switch fromVal + toVal {
        case CELSIUS + CELSIUS, FAHRENHEIT + FAHRENHEIT, KELVIN + KELVIN:
            for id, val := range rawArgs {
                vall, _ := strconv.ParseFloat(val, 64)
                outputValues[id] = vall
            }
        case CELSIUS + FAHRENHEIT:
            for id, val := range convValues {
                outputValues[id] = float64(tempconv.CtoF(tempconv.Celsius(val)))
            }
        case CELSIUS + KELVIN:
            for id, val := range convValues {
                outputValues[id] = float64(tempconv.CtoK(tempconv.Celsius(val)))
            }
        case FAHRENHEIT + CELSIUS:
            for id, val := range convValues {
                outputValues[id] = float64(tempconv.FtoC(tempconv.Fahrenheit(val)))
            }
        case FAHRENHEIT + KELVIN:
            for id, val := range convValues {
                outputValues[id] = float64(tempconv.FtoK(tempconv.Fahrenheit(val)))
            }
        case KELVIN + CELSIUS:
            for id, val := range convValues {
                outputValues[id] = float64(tempconv.KtoC(tempconv.Kelvin(val)))
            }
        case KELVIN + FAHRENHEIT:
            for id, val := range convValues {
                outputValues[id] = float64(tempconv.KtoF(tempconv.Kelvin(val)))
            }
        default:
            fmt.Printf(tempUnitERR, fromVal)
            os.Exit(1)
        }

    // TODO: other conversion packages can be used here
    case LENGTH:
    case WEIGTH:
    default:
        fmt.Printf(typeERR, convType)
        os.Exit(1)
    }

    // return results to os.Stdout
    var rawOutput string
    for _, val := range outputValues {
        rawOutput += strconv.FormatFloat(val, 'f', 2, 64) + " "
    }
    if _, err := fmt.Fprint(os.Stdout, rawOutput); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
