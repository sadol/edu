package main

import (
	"./page"
	"fmt"
)

func main() {
	firstBody := []byte{'f', 'i', 'r', 's', 't'}
	pageOne := page.NewPage("first_batch", firstBody...)
	pageTwo := page.NewPage("second_batch")
	fmt.Println()
	fmt.Println("Originals:")
	fmt.Println()
	fmt.Printf("%v\n", pageOne)
	fmt.Printf("%v\n", pageTwo)
	pageOne.Save()
	pageTwo.Save()
	fmt.Println()
	fmt.Println()
	fmt.Println("Copies:")
	fmt.Println()
	if pageOneCopy, err := page.Load(pageOne.GetName()); err == nil {
		fmt.Println(pageOneCopy)
	} else {
		panic(err)
	}
	if pageTwoCopy, err := page.Load(pageTwo.GetName()); err == nil {
		fmt.Println(pageTwoCopy)
	} else {
		panic(err)
	}
}
