// using external custom package

package main

import (
	"./greetings" // relative path in this case
	"fmt"
)

func main() {
	if greetings.IsAm() {
		fmt.Println(greetings.GoodDay)
	} else {
		fmt.Println(greetings.GoodNight)
	}
}
