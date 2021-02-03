/*
Tagging variables and anonymous variables in golang.
Anonymous vars are used in `golang inheritance' (that means composition).
Tagging is used with `reflect' package.
*/
package main

import (
    "fmt"
    "reflect"
)

type myStruct struct {
	f1     float64 "named field"
	int    "anonymous field int - only 1 such thing in a structure"
	string "anonymous field  - only 1 such thing in a structure"
}

func main() {
    // creating instance of the struct using literal expression
	instance1 := myStruct{34.5, 4, "ala"}
    msType := reflect.TypeOf(instance1)                       // reflect object
    tagField := msType.Field(2)                          // extracted 2nd field
    fmt.Println(tagField.Tag)                   // extracted tag from 2nd field
	fmt.Println(instance1)
}
