// callback to replace all non ascii chars with ? character

package main

import (
	"fmt"
	"strings"
)

func main() {
	testingString := "Łękołody"
	processedString := strings.Map(replaceNonAscii, testingString)
	fmt.Printf("Under test : %s\t\t, processed : %s.\n",
		testingString, processedString)
}

// callback function
func replaceNonAscii(c rune) rune {
	if c > 255 {
		return '?'
	}
	return c
}
