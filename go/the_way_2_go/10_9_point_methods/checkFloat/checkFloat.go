package checkFloat

import (
	"../array"
	"math"
)

func CheckMultOverflow(operands ...float64) bool {                  // variadic
	_, bigest, noOfElements := array.MinMaxArrFloat64(operands...)

	if noOfElements == 1 {             // no need to asses for only one element
		return false
	}
	// veeeeery crude
	if math.IsInf(math.Pow(bigest, float64(noOfElements)), 1) {
		return false
	}
	return true
}
