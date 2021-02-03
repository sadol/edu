// fun with recursive function

package main

import "fmt"

func main() {
	max := 10
	printrec(max)
}

func printrec(i int) {
	fmt.Println(i)
	if i <= 1 {                          // exit clause from recursive function
		return
	}
	printrec(i - 1)
}
