package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const EOT = 'S' // end of transmission sign

func main() {
	inputReader := bufio.NewReader(os.Stdin) // new reader to read from stdin
	fmt.Println("Please enter some text (`S' to exit) → ")
	inputFromStdin, err := inputReader.ReadString(EOT) // here program reads
	if err != nil {
		fmt.Println("Read error, exiting program.")
		return
	}
	// `inputFromStdin' contains terminating character also
	fmt.Printf("Number of chars (new lines & EOT excluded): %d.\n",
		howManyChars(inputFromStdin))
	fmt.Printf("Number of words: %d.\n", howManyWords(inputFromStdin))
	fmt.Printf("Number of lines: %d.\n", howManyLines(inputFromStdin))
}

func howManyChars(input string) int {
	newLines, allChars := howManyLines(input), len(input)
	if newLines == 0 && allChars > 1 { // only one line terminater with EOT
		newLines++
	}
	return allChars - newLines
}

func howManyWords(input string) int {
	words := strings.Fields(input)
	noOfWords := len(words)
	for _, value := range words { // remove sole EOT word if exists
		if value == string(EOT) {
			noOfWords--
		}
	}
	return noOfWords
}

func howManyLines(input string) int {
	// last line always ends with EOT (not EOL)
	return strings.Count(input, "\n") + 1
}
