// fun with switch construct continues

package main

import "fmt"

func main() {
	message := ""
	for i := 1; i <= 100; i++ {
		switch {         // very usefull form of switch, better than if else if
		case i%15 == 0:                // least probable case first (this time)
			message = "FizzBuzz"
		case i%5 == 0:
			message = "Buzz"
		case i%3 == 0:                   // most probable case last (this time)
			message = "Fizz"
		default:
			message = fmt.Sprintf("%d", i)
		}
		fmt.Println(message)
	}
}
