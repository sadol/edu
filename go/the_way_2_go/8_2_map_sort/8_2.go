// sorting maps

package main

import (
	"fmt"
	"sort"
)

func main() {
    polEngDict := map[string]string{"krzesło" : "chair",
                                    "masło" : "butter",
                                    "hasło" : "password"}

	for key, value := range polEngDict {
		fmt.Printf("pol: <%s>\t eng: <%s>\n", key, value)
	}

	fmt.Println()
    fmt.Println("Sorted pol->eng:")
	printSortedMap(polEngDict)

	//inverting map procedure (key-value relationship must be unique)
	engPolDict := make(map[string]string, len(polEngDict))
	for key, value := range polEngDict {
		engPolDict[value] = key
	}

	fmt.Println()
    fmt.Println("Sorted eng->pol:")
	printSortedMap(engPolDict)
}

// sorting keys procedure for maps:
func printSortedMap(dict map[string]string) {
	i, sortedKeys := 0, make([]string, len(dict)) // 1. create slice of keys
	for key, _ := range dict {                    // 2. populate it
		sortedKeys[i] = key
		i++
	}

	sort.Strings(sortedKeys)                      // 3. sort it
	for _, value := range sortedKeys {
		fmt.Printf("<%s>\t<%s>\n", value, dict[value])
	}
}
