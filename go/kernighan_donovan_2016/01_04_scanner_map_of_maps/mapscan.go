// excercise 1.4, counting unique rows using maps and scanners.
// Printing file names of the duplicated lines.
// Test usage: `mapscan first.txt second.txt first.txt'
package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    counts := make(map[string]map[string]int)                    // map of maps
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            if f, err := os.Open(arg); err != nil {
                fmt.Fprintf(os.Stdout, "dup2: %v.\n", err)
                continue
            } else {
                countLines(f, counts)
                f.Close()
            }
        }
    }

    for fileKey, dupMap := range counts {
        for lineKey, duplicates := range dupMap {
            if dupMap[lineKey] > 1 {
                fmt.Printf("%s:\t< %s >\t duplicates: %d.\n", fileKey , lineKey, duplicates)
            }
        }
    }
}

// INFO: no error checking!
func countLines(f *os.File, counts map[string]map[string]int) {
    var fileName string = f.Name()
    input := bufio.NewScanner(f)
    if _, present := counts[fileName]; !present { // filename key not found
        counts[fileName] = make(map[string]int)
    } else {                      // to avoid double reading from the same file
        return
    }

    for input.Scan() {
        counts[fileName][input.Text()]++          // very nice golang map idiom
    }
}
