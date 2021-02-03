/*
	Check if integer division by 0 can crash program (causing runtime panic).
*/

package main

func main() {
	var first int = 10
	second := 0
	result := first / second
	print(result)
}
