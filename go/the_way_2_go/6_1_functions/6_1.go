// Function which accepts 2 ints & returns their sum, product & diff.

package main

import "fmt"

func main() {
	first, second := 100, 30
	sum, product, difference := spd_named(first, second)
	fmt.Printf("Spd_named results=> sum=%d\tproduct=%d\tdifference=%d\n",
		sum, product, difference)
	sum, product, difference = spd_unnamed(first, second)
	fmt.Printf("Spd_unnamed results=> sum=%d\tproduct=%d\tdifference=%d\n",
		sum, product, difference)
}

// named return variables (very useful & consistent)
func spd_named(first int, second int) (sum int, product int, difference int) {
	sum = first + second
	product = first * second
	difference = first - second
	return
}

// non named return vars version
func spd_unnamed(first int, second int) (int, int, int) {
	return first + second, first * second, first - second
}
