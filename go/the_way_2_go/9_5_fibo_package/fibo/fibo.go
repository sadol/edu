// helper package file

package fibo

import "math"

// generator function -> only last state is remembered, there is no need
// to engage global variables if there is closure at hand
func Fibo(operation string) func() (int64, int) {
	var Prev int64 = 1
	var Last int64 = 2
	var LastIndex = 2
	var isOverflowed bool = false
	var maxPrev int64 = 0
	var f func(int64, int64) int64

    // setting overflow thresholds for particular functors
	switch operation {
	case "+":
		maxPrev = int64(math.Round(0.5 * math.MaxInt64))
		f = add
	case "*":
		maxPrev = int64(math.Round(0.1 * math.MaxInt64))
		f = multiple
	default:
		maxPrev = 0  // it's not even funny anymore
		f = trump
	}

	return func() (int64, int) {
		if isOverflowed {
			return 0, LastIndex
		} else {
			if Prev > maxPrev { // overflow fuse
				isOverflowed = true
				return 0, LastIndex
			}
			Prev, Last = Last, f(Prev, Last)
			LastIndex++
			return Last, LastIndex
		}
	}
}


//`private' functions

func add(operand1 int64, operand2 int64) int64 {
	return operand1 + operand2
}

func multiple(operand1 int64, operand2 int64) int64 {
	return operand1 * operand2
}

// fake function
func trump(operand1 int64, operand2 int64) int64 {
	return 0 // of course it should return 0, didn't it?
}
