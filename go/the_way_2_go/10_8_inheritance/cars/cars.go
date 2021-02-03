// simple cars package

package cars

import (
    "../engine"
    "fmt"
)

//INFO: public interfaces must start with CAPITAL LETTER!!!
type Car struct {
	engine.Engine     "anonymous element of the struct"
	wheelCount uint "this is NOT exported; accsess through getter and setter only"
}

//getter
func (car *Car) NoWheels() uint {
	return car.wheelCount
}

//setter
func (car *Car) SetNoWheels(howMuch uint) {
	car.wheelCount = howMuch
}

type Mercedes struct {
	Car "and now Merc stuct can use methods of embedded Car struct directly"
}

func (merc *Mercedes) SayHiToMerkel() {
	fmt.Println("Hallo from Berlin.")
}
