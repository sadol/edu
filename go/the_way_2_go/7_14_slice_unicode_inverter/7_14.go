// Reversing in golang.

package main

import "fmt"

func main() {
	testedString := "Łękołody"
	fmt.Printf("Original string: <%s>, reversed: <%s>.\n",
		testedString, Reverse(testedString))
}

// reverser
func Reverse(inputString string) string {
	reversedString := make([]rune, len(inputString))
	for index, char := range inputString {
		reversedString[len(reversedString)-index-1] = rune(char)
	}
	return string(reversedString)
}
