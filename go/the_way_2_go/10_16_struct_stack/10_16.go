package main

import (
	"./arrstack"
	"fmt"
)

func main() {
	const LIMIT = 4
	myStack := arrstack.NewStack(LIMIT)

	fmt.Printf("%v.\n", myStack)
	fmt.Println()
	fmt.Println("Filling the stack.")
	for i := 0; i < LIMIT; i++ {
		if myStack.Push(10 + i) {
			fmt.Printf("%v.\n", myStack)
		}
	}
	fmt.Println("Filled stack.")
	fmt.Printf("%v.\n", myStack)
	fmt.Println()
	fmt.Println("Clearing the stack.")
	for i := LIMIT; i >= 0; i-- {
		if element, ok := myStack.Pop(); ok {
			fmt.Printf("%v\tpopped element: %d.\n", myStack, element)
		}
	}
	fmt.Println("Cleared stack.")
	fmt.Printf("%v.\n", myStack)

}
