// using := for error handling

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	orig := "ABC"
	converted, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Println(err)             // try to put some meanigful error message
		os.Exit(1)                    // error ready to be catched by the shell
	}
	fmt.Printf("The integer is %d\n", converted)
}
