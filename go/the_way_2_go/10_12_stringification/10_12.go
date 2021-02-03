// stringification in golang

package main

import (
	"fmt"
	"strconv"
)

//-------------------------T----------------------
type T struct {
	a int
	b float32
	c string
}

// stringification of struct, INFO: do not use `fmt' to return from String()
// functions because of risk of infinite call loop `fmt'→String()→`fmt'→...
func (t *T) String() string {
	return "" + strconv.Itoa(t.a) + " / " + strconv.FormatFloat(float64(t.b), 'f', 6, 32) +
		" / " + strconv.Quote(t.c)
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("`T' stringification: <%v>.\n", t)
}
