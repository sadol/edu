package main

import (
    "fmt"
)

// removes adjacent duplicates from a slice of strings (memory waste?)
func removeDup(input []string) []string {
    temp := make([]string, 0)
    for id, val := range input {
        if id < len(input) - 1 {
            if val == input[id + 1] { continue }
        }
        temp = append(temp, val)
    }
    return temp
}

func main() {
    test1 := []string{"Putin", "chuj", "chuj", "Xi", "pała", "pała", "pała",
                      "Orban", "dupa", "dupa"}
    fmt.Printf("\ninitial slice: %v.\n", test1)
    test1 = removeDup(test1)
    fmt.Printf("duplicates filtered slice: %v.\n\n", test1)
}
