// variadic function intro

package main

import "fmt"

func main() {
	ints := []int{3, 4, 5, 56, -56}
	ints0 := []int{}
	fmt.Println("Variadic function output(populated table):")
	varfun(ints...)
	fmt.Println()
	fmt.Println("Variadic function output(empty table):")
	varfun(ints0...)
}

func varfun(intArr ...int) {
	if len(intArr) == 0 {
		fmt.Println("Nope.") // like in BASH
		return
	}
	for _, element := range intArr {
		fmt.Println(element)
	}
}
