//

package main

import (
	"./cars"
	"fmt"
)

func main() {
	myMerc := new(cars.Mercedes)
	myMerc.SetNoWheels(uint(5))
	myMerc.SetNoCylinders(uint(5))
	fmt.Printf("How many wheels myMerc has: %d.\n", myMerc.NoWheels())
	fmt.Printf("How many cylinders myMerc has: %d.\n", myMerc.NoCylinders())
	fmt.Println()
	fmt.Println("myMerc starts:")
	myMerc.Starts()
	fmt.Println()
	fmt.Println("My merc talks!!!")
	myMerc.SayHiToMerkel()
	fmt.Println()
	fmt.Println("myMerc stops:")
	myMerc.Stops()
}
