package main

import (
	"./bookreader"
	"fmt"
)

func main() {
	fileName := "products.txt"
	var myBooks bookreader.Books
	var bookError error

	if myBooks, bookError = bookreader.LoadBooks(fileName); bookError != nil {
		fmt.Println("Oops")
	} else {
		for _, element := range myBooks {
			fmt.Println(&element)
		}
	}
}
