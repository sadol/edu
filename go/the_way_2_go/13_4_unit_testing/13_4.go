package main

import (
	"./reverse"
	"fmt"
)

func main() {
	testedString := "Łękołody"
	fmt.Printf("Original string: <%s>, reversed: <%s>.\n",
		testedString, reverse.Reverse(testedString))
}
