// stringification of int alias

package main

import "fmt"

//---------------------TIMEZONE-----------------
type TZ int

const (
	UTC TZ = iota
	WEU
	EEU
	CAI
)

func (t TZ) String() string {
	zones := map[TZ]string{
		UTC: "Universal Greenwich Time",
		WEU: "West European Time",    // bogus
		EEU: "Eastern European Time", // bogus
		CAI: "Cairo Time",            // bogus
	}

	return zones[t]
}

func main() {
	fmt.Printf("%v.\n", UTC)
	fmt.Printf("%v.\n", WEU)
	fmt.Printf("%v.\n", EEU)
	fmt.Printf("%v.\n", CAI)
}
