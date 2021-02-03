package engine

import "fmt"

//INFO: public interfaces must start with CAPITAL LETTER!!!
type Engine struct {
	noOfCylinders uint "unexported number of cylinders"
}

// Engine setter
func (eng *Engine) SetNoCylinders(howMuch uint) {
	eng.noOfCylinders = howMuch
}

//Engine getter (even getter is called on behalf of a POINTER to avoid memory
//overburden)
func (eng *Engine) NoCylinders() uint {
	return eng.noOfCylinders
}

func (eng *Engine) Starts() {
	fmt.Println("Engine starts")
}

func (eng *Engine) Stops() {
	fmt.Println("Engine stops")
}
