package main

import (
	"./rsimple"
	"./simple"
	"./simpler"
	"fmt"
)

func main() {
	var myInterfaces [2]simpler.Simpler
	newVal := 10
	myInterfaces[0] = simple.NewSimple(5)
	myInterfaces[1] = rsimple.NewRSimple(10)
	fmt.Printf("%v.\n", myInterfaces[0])              // String() interface test
	fmt.Printf("Value: %d.\n", myInterfaces[0].Get()) // Get() interface test
	myInterfaces[0].Set(newVal)                       // Set() interface test
	fmt.Printf("%v.\n", myInterfaces[0])
	fmt.Printf("%v.\n", myInterfaces[1])              // String() interface test
	fmt.Printf("Value: %d.\n", myInterfaces[1].Get()) // Get() interface test
	newVal = 20
	myInterfaces[1].Set(newVal) // Set() interface test
	fmt.Printf("%v.\n", myInterfaces[1])

	fmt.Println()
	// i can't use : fI(myInterfaces...) nor fI(myInterfaces)
    fI(myInterfaces[0], myInterfaces[1])
}

func fI(interfaces...interface{}) {
	for _, value := range interfaces {
        switch value.(type) { // golang idiom for checking types using interface
        case *simple.Simple:
		    fmt.Println("Type is *Simple")
        case *rsimple.RSimple:
		    fmt.Println("Type is *RSimple")
        default:
            fmt.Printf("Unknown type %T.\n", value)
        }
	}
}
