package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("first.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()
	outputFile, err := os.OpenFile("first_out.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
    i := 1
    for {
		inputString, _, err := inputReader.ReadLine()
		if err == io.EOF {
			fmt.Println("Conversion done! Check output file.")
			return
		}
		outputString := string(inputString[2:5]) + "\n"
		if _, err := outputWriter.WriteString(outputString); err != nil {
			fmt.Println(err)
			return
		} else {
			outputWriter.Flush() // this shoud NEVER be skipped !!!
		}
	}
}
