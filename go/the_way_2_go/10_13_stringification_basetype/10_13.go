// stringification in golang

package main

import (
	"fmt"
	"strconv"
)

//----------------------CELSIUS-------------------------
type Celsius float64

//WARNING: do not put asterix in the reciever part of function in functions
//		   which expand functionality of BASE types of GO (int, float etc.)!!!
//		   (c *Celsius)... → this is wrong because this implies explicit usage
//		   of String() interface!!!
func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "\u2070C"
}

func main() {
	var temp1 Celsius = 23.45
	var temp2 Celsius = -34.0005
	fmt.Printf("First temperature : %v.\n", temp1)
	fmt.Printf("Second temperature : %v.\n", temp2)
}
