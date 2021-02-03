package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := ""
    separator := ", "
	if len(os.Args[1:]) > 0 {
	    who += strings.Join(os.Args[1:], separator)
	}
	fmt.Printf("Hello %s!", who)
}
