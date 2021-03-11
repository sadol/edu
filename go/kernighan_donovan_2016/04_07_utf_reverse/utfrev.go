package main

import (
    "fmt"
    "strings"
)

// UTF-8 reverser
func reverseUTF(input []byte) []byte {
    reader := strings.NewReader(string(input))
    temp := make([]byte, len(input))
    var size int
    var run rune
    var err error

    for startByte := len(input); reader.Len() > 0; startByte -= size {
        if run, size, err = reader.ReadRune(); err == nil {
            copy(temp[startByte - size:startByte], string(run))
        } else { break }
    }

    copy(input, temp)
    return input
}

func main() {
    test1 := "Putin chuj złamany."
    fmt.Println("ORIGINAL: ", test1)
    test1 = string(reverseUTF([]byte(test1)))
    fmt.Println("REVERSED: ", test1)
}
