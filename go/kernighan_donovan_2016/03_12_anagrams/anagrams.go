package main

import "fmt"

// creates map of characters of some string
func createMap(someString string) (output map[string]int) {
    output = make(map[string]int)
    for i := 0; i < len(someString); i++ {
        output[string(someString[i])]++
    }
    return output
}

// checks anagrams
func areAnagrams(first, second string) bool {
    if len(first) != len(second) || first == second || len(first) == 0 { return false }
    mapFirst := createMap(first)                               // maps of chars
    mapSecond := createMap(second)

    for id, val := range mapFirst {
        if mapSecond[id] != val { return false }
    }
    return true
}

func main() {
    tests := [][]string{{"ala", "ola"}, {"zupa", "uzap"}, {"astrolabium", "astrobalium"}, {"kook", "okko"}}
    fmt.Println()
    for _, val := range tests {
        fmt.Printf("first: <%s>\tsecond: <%s>\tanagram? : <%s> \n", val[0], val[1], areAnagrams(val[0], val[1]))
    }
    fmt.Println()
}
