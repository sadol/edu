// recovering from run time grave error
package main

import (
	"fmt"
	"log"
)

type int2function func(first int, second int) int // to simplify other declarations

func divider(first, second int) int { // some function to test (compatible with type def above)
	return first / second // panics here in case of division by zero
}

// INFO: not very generic function (should be implemented with interface{} &
//       type assertions).
func errorHandler(myFunction int2function) int2function { // CLOSURE of multipurpose function type over the actual args supplied by the caller
	return func(f, s int) int { // caller is responsible for supplying proper args
		defer func() { // defered function to check panics
			if err, ok := recover().(error); ok { // panics checks by using builin RECOVER + type assertion
				log.Printf("Controlled run time panic: %v.", err) // GRACEFUL solve of panic (I'm in full controll, not the operating system)
			}
		}() // defered function must be called! (compare to `defer fileName.Close()')
		return myFunction(f, s) // function call is implemented indirecly, output must be catch by the caller
	}
}

func main() {
	testCases := [][2]int{{1, 2}, {4, 2}, {0, 4}, {4, 0}}
	funcToCall := errorHandler(divider) // catch function `pointer', one for all calls in this case
	for _, pair := range testCases {
		returnValue := funcToCall(pair[0], pair[1])                      // use this pointer
		fmt.Printf("divide on args: %v, gives %v.\n", pair, returnValue) // use output if it's correct
	}
}
