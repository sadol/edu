package main

import (
    "fmt"
    "unicode"                                                        // IsSpace
)

func squashInplace(input []byte) []byte {
    var consecutiveWhites int
    var allWhites int
    initLen := len(input)

    for id, val := range input {
        if unicode.IsSpace(rune(val)) {       // is there a problem with \v ???
            consecutiveWhites++
            continue
        }

        switch {
        case consecutiveWhites == 1:
            consecutiveWhites = 0
        case consecutiveWhites > 1:
            // one white is only one byte?
            allWhites += consecutiveWhites
            copy(input[id - consecutiveWhites + 1:], input[id:])
            consecutiveWhites = 0
        }
    }
    return input[:initLen - allWhites]
}

func squash(input []byte) []byte {
    var consecutiveWhites int
    temp := make([]byte, 0)

    for _, val := range input {
        if unicode.IsSpace(rune(val)) {
            consecutiveWhites++
            continue
        }

        switch {
        case consecutiveWhites != 0:
            consecutiveWhites = 0
            temp = append(temp, ' ')
            fallthrough
        default:
            temp = append(temp, val)
        }
    }

    input = make([]byte, len(temp))
    copy(input, temp)
    return input
}

func main() {
    test1 := "A Putin  \t to  \v  chuj \n  złamany.   "
    fmt.Println("-------------------------")
    fmt.Println("BEFORE:")
    fmt.Println([]byte(test1))
    fmt.Println(test1)
    test1 = string(squash([]byte(test1)))
    fmt.Println("-------------------------")
    fmt.Println("AFTER:")
    fmt.Println([]byte(test1))
    fmt.Println(test1)
}
