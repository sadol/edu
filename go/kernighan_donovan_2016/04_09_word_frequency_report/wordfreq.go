package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

func main() {
    words := make(map[string]int)                    // counter of words
    var noWords int                                  // number of words
    freq := make(map[string]float64)                 // frequency of words
    var keys []string                                // for sorting map keys
    in := bufio.NewReader(os.Stdin)
    scan := bufio.NewScanner(in)

    scan.Split(bufio.ScanWords)

    for scan.Scan() {
        words[scan.Text()]++
        noWords++
    }

    if err := scan.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "scanner error: ", err)
    }

    if noWords == 0 { os.Exit(0) }

    for word, count := range words {
        freq[word] = float64(count) / float64(noWords)
    }

    for word := range words {                        // search in map keys
        keys = append(keys, word)
    }
    sort.Strings(keys)

    fmt.Println("WORD\tFREQUENCY\n")
    for _, key := range keys {                      // "sorted" map keys usage
        fmt.Printf("%s\t%4.2f %%\n", key, freq[key] * 100.0)
    }
}
