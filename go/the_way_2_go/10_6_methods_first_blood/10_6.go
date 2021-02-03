/*
Methods in golang (traditional `OBJECT.METHOD(ARGS)' notation).
*/

package main

import (
    "fmt"
    "./employee"
)

func main() {
	emp1 := new(employee.Employee)
	emp1.Salary = 6000.00
	fmt.Printf("Salary before raise: %.2f.\n", emp1.Salary)
	emp1.GiveRaise(50.04)
	fmt.Printf("Salary after raise: %.2f.\n", emp1.Salary)
}
