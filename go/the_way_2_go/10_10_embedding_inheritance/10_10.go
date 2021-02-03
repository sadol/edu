// golang inheritance by embedding anonyomus substructs & overshadowing methods

package main

import (
    "./employee"
)

func main() {
	emp1 := employee.NewEmployee(111, "Jan", "Novak", 6000)
	emp1.PrintEmployee()
}
