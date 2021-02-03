// Split testing

package main

import "fmt"

func main() {
	evenTestString := "Fuck Google!"
    oddTestString := "Fuck Facebook:)"
	fmt.Printf("Tested string: <%s>, first part: <%s>, second part: <%s>.\n",
		evenTestString, evenTestString[:len(evenTestString)/2],
		evenTestString[len(evenTestString)/2:])
	fmt.Printf("Tested string: <%s>, first part: <%s>, second part: <%s>.\n",
		oddTestString, oddTestString[:len(oddTestString)/2],
		oddTestString[len(oddTestString)/2:])
}
