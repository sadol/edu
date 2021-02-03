// Splitting strings to two substrings

package main

import "fmt"

func main() {
	position, testedString := 3, "Ala ma kota, a kot ma Alę."
	fmt.Printf("String under testing: <%s>.\n", testedString)
	first, second := Split(testedString, position)
	fmt.Printf("After split at position %d, first substring: <%s>", position, first)
	fmt.Printf(", second substring: <%s>.\n", second)
}

// splitter function
func Split(stringToSplit string, position int) (first, second string) {
	first, second = stringToSplit[:position], stringToSplit[position:]
	return
}
