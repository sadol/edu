package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    counts := make(map[rune]int)                    // counter of unicode chars
    letters := make(map[rune]int)                   // counter of letters
    var digits [10]int                              // counter of digits
    var utflen [utf8.UTFMax + 1]int                 // counter of lengths of UTF encodings
    invalid := 0                                    // counter of the invalid characters

    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune()
        if err == io.EOF { break }                  // loop normal exit point
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount %v.\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        counts[r]++                                 // nice golang map idiom
        utflen[n]++
        if unicode.IsDigit(r) { digits[r]++ }
        if unicode.IsLetter(r) { letters[r]++ }
        // ... etc with other `unicode tests' ...
    }

    fmt.Printf("rune\tcount\n")
    for c, n := range counts { fmt.Printf("%q\t%d\n", c, n) }
    fmt.Printf("len\tcount\n")
    for i, n := range utflen {
        if i > 0 { fmt.Printf("%d\t%d\n", i, n) }
    }
    fmt.Printf("digit\tcount\n")
    for d, n := range digits { fmt.Printf("%d\t%d\n", d, n) }
    fmt.Printf("letter\tcount\n")
    for l, n := range letters { fmt.Printf("%q\t%d\n", l, n) }
    // ... etc with other `unicode tests' ...
    if invalid > 0 {
        fmt.Println("\ninvalid chars: %d.\n", invalid)
    }
}
