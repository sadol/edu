// Playing with uniqueness.

package main

import "fmt"

func main() {
	test := "Ala ma kota."
	fmt.Printf("String to test: <%s>, unique runes: <%s>.\n", test, Uniq(test))
	test = "Coffevee."
	fmt.Printf("String to test: <%s>, unique runes: <%s>.\n", test, Uniq(test))
}

// Traverses array of chars and copies them to another only if one char is
// different than its predecessor.
func Uniq(input string) string {
	output := make([]rune, 1)
	output[0] = rune(input[0])
    for i := 1; i < len(input); i++ {
        if input[i] != input[i-1] {
            output = append(output, rune(input[i]))
        }
	}
	return string(output)
}
