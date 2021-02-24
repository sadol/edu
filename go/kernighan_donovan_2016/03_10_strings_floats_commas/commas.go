package main

import (
    "fmt"
    "strconv"
    "strings"
)

// inserts commas in proper places of the string representation of float
func floatingCommas(input string) (output string, err error) {
    if _, err = strconv.ParseFloat(input, 64); err != nil { return }
    dotPosition := strings.Index(input, ".")
    j := 1
    if dotPosition != -1 {
        if len(input[dotPosition:]) <= 3 && len(input[:dotPosition]) <= 3 {
            return input, nil
        }
        output = "."
        for i := dotPosition - 1; i >= 0; i-- {
            output = string(input[i]) + output
            if j % 3 == 0 { output = "," + output }
            j++
        }
        j = 1
        for i := dotPosition + 1; i < len(input); i++ {
            output += string(input[i])
            if j % 3 == 0 { output += "," }
            j++
        }
    } else {
        if len(input) <= 3 { return input, nil }
        j = 1
        for i := len(input) - 1; i >= 0; i-- {
            output = string(input[i]) + output
            if j % 3 == 0 { output = "," + output }
            j++
        }
    }
    if string(output[0]) == "," { output = output[1:] }
    if string(output[len(output) - 1]) == "," { output = output[:len(output) - 1] }
    return
}

func main() {
    data := []string{"ala", "12.g34", "456", "12342.456745678", "34.45",
                     "1001.1001", "1234", "123.1234"}
    fmt.Println()
    for id, val := range data {
        out, err := floatingCommas(val)
        fmt.Printf("%d:\toriginal: <%s>\terr: <%s>\tprocesed: <%s>.\n",
                   id, val, err, out)
    }
    fmt.Println()
}
