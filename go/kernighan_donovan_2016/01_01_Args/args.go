// exercise 1.1 - add Args[0] to the main program
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[:], " "))
}
