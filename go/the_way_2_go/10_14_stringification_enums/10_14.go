// stringification of int alias

package main

import "fmt"

//-------------------DAY---------------------------
// golang enum
type Day int

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

func (d Day) String() string {
	daynames := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday",
		"Friday", "Saturday", "Sunday"}
	return daynames[d]
}

//-------------------------------------------------

func main() {

	fmt.Printf("%v.\n", MO)
	fmt.Printf("%v.\n", Day(4))
}
