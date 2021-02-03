package main

import (
	"./genstack"
	"fmt"
)

func main() {
    // stack is operated by generic interface
    myStack := genstack.NewStack()
    ints := genstack.Vector{1, 2, 3, 4, 5, 6}
    floats := genstack.Vector{3.14, 3.33, 7.77}
    strings := genstack.Vector{"ala", "ola", "basia"}
    varia := genstack.Vector{"ala", 1, 1.11}

	fmt.Printf("%v.\n", myStack)
	fmt.Println()
	fmt.Println("Filling the stack.")
	for _, value := range ints {
		myStack.Push(value)
	}
	for _, value := range floats {
		myStack.Push(value)
	}
	for _, value := range strings {
		myStack.Push(value)
	}
	for _, value := range varia {
		myStack.Push(value)
	}
	fmt.Println("Filled stack.")
	fmt.Printf("%v.\n", myStack)
	fmt.Println()
	fmt.Println("Clearing the stack.")
	for i := myStack.Len(); i > 0; i-- {
		myStack.Pop()
	}
	fmt.Println("Cleared stack.")
	fmt.Printf("%v.\n", myStack)
}
