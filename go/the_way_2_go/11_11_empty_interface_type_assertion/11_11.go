package main

import (
	"fmt"
)

type Element interface{}

func main() {
	var test1 = []Element{1, 2, 3, 4, 5, 6, 7, 8}
	var test2 = []Element{"ala", "ola", "basia"}
	var test3 = []Element{1.99, 3.88, 4.77}
	ok1 := genericMapFunc(genericLambda2, test1...)
	ok2 := genericMapFunc(genericLambda2, test2...)
	ok3 := genericMapFunc(genericLambda2, test3...)
	ok4 := genericMapFunc(genericLambda2)
	if ok1 {
		fmt.Println("Mapped element:", test1, ".")
	}
	if ok2 {
		fmt.Println("Mapped element:", test2, ".")
	}
	if ok3 {
		fmt.Println("Mapped element:", test3, ".")
	}
	if ok4 {
		fmt.Println("OK 4") // impossibru
	}
}

// generic versions of the mapping function: do not use specific structs as
// actual arguments, use generic interfaces instead
func genericMapFunc(f func(Element) Element, input ...Element) (ok bool) {
	if len(input) > 0 {
		switch input[0].(type) { // type SWITCH
		case int, string:
			for index, value := range input {
				input[index] = f(value)
			}
			ok = true
		default:
			ok = false
		}
	} else {
		ok = false
	}
	return
}

// generic version of the lambda * 10 function
func genericLambda2(any Element) (output Element) {
	switch any.(type) {
	case int:
		output = any.(int) * 2 // type ASSERTION is needed because of operator * which is not defined in Element
	case string:
		output = any.(string) + any.(string) // type assertion is needed because of + operator which is not defined in Element
	default:
		output = nil
	}
	return
}
