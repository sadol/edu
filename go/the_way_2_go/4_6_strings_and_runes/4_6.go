// counting characters (runes) and bytes

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"                                         // for rune handling
)

func main() {
	asciiString := "Basia Asia Joasia Jasia"                // ASCII only chars
	headerFormat := "%-3s\t%-5s\t%-5s\t%-5s\n"
	dataFormat := "%-3d:\t%-5d\t%-5d\t%-5c\n"
	utf8String := "Basia Asia Jadźia Józia"                  // NON ASCII chars
	fmt.Println()
	fmt.Println("Number of characters in string is:", len(asciiString),
		"and number of bytes is", len(asciiString))
	fmt.Println("Byte by byte print of the `asciiString':")
	fmt.Println()
	fmt.Printf(headerFormat, "No.", "Bytes", "Code", "Char")
	fmt.Println("-----------------------------")
	for b := 0; b < len(asciiString); b++ {
		fmt.Printf(dataFormat, b, 1, asciiString[b], asciiString[b])
	}

	fmt.Println()
	fmt.Println("Number of characters in string is:", utf8.RuneCountInString(utf8String),
		"and number of bytes is", len(utf8String))
	fmt.Println("Byte by byte print of the `utf8String':")
	fmt.Printf(headerFormat, "No.", "Bytes", "Code", "Char")
	fmt.Println("----------------------------")
	var myReader *strings.Reader = strings.NewReader(utf8String)
	var myRune rune
	var mySize int
	var counter int = 0                                   // not very beautiful
	for myReader.Len() > 0 {
		myRune, mySize, _ = myReader.ReadRune()
		fmt.Printf(dataFormat, counter, mySize, myRune, myRune)
		counter++
	}
}
