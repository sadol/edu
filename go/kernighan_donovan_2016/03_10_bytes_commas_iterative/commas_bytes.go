// Non-recursive version of the comma insertion algorithm.
package main

import (
    "fmt"
    "bytes"
    "strconv"
    "strings"
    "errors"
)

func commaBytes(input string) (output string, err error) {
    // 1. remove white chars:
    subInput := strings.ReplaceAll(input, " ", "")
    subInput = strings.ReplaceAll(subInput, "\n", "")
    subInput = strings.ReplaceAll(subInput, "\t", "")
    // 2. check if `input' is a legal number chars string
    if _, err = strconv.ParseFloat(subInput, 64); err != nil { return }
    if strings.ContainsAny(subInput, "aAbBcCdDeEfF") {
        return "", errors.New("Scientific notation or hex floats detected.")
    }
    // 3. check for minus sign(plus sign)
    var outputBuf bytes.Buffer
    if strings.Contains(subInput, "-") {
        subInput = strings.ReplaceAll(subInput, "-", "")
        outputBuf.WriteByte('-')
    } else {
        subInput = strings.ReplaceAll(subInput, "+", "")
    }

    inputBuf := *bytes.NewBufferString(subInput)
    commaIndex := inputBuf.Len() % 3
    if commaIndex == 0 { commaIndex = 3 }

    for i := 0; i < inputBuf.Len(); i++ {
        if i == commaIndex {
            outputBuf.WriteByte(',')
            commaIndex += 3
        }
        outputBuf.WriteByte(inputBuf.Bytes()[i])
    }

    return outputBuf.String(), nil
}

func main() {
    data := []string{"ala", "12.g34", "456", "12342.456745678", "34.45",
                     "1001.1001", "1234", "123.1234", ".01", ".001", ".0001",
                     "1000.", "1.", "100.", "AB10", "ab12", ".23e10", "-1",
                     "-1000.01", "-1000.0001"}
    fmt.Println()
    for id, val := range data {
        out, err := commaBytes(val)
        fmt.Printf("%d:\toriginal: <%s>\terr: <%s>\tprocesed: <%s>.\n",
            id, val, err, out)
    }
    fmt.Println()

}
