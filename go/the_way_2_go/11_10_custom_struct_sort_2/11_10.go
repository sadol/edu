package main

import (
	"./float64array"
	"./miner"
	"fmt"
)

func main() {
	const size = 10
	var example1 = float64array.NewFloat64Array()
	var example2 = float64array.NewFloat64Array(size)
	var minIndex1, minIndex2 int

	fmt.Printf("%v\n", example1)
	fmt.Printf("%v\n", example2)
    // float64Array is implicitly `Miner' interface also
	minIndex1, minIndex2 = miner.Min(example1), miner.Min(example2)
	fmt.Printf("Minimal values → minValue1 = %f, minValue2 = %f.\n",
		example1.Get(minIndex1), example2.Get(minIndex2))
	fmt.Printf("Minimal indexes → minIndex1 = %d, minIndex2 = %d.\n", minIndex1,
		minIndex2)
}
