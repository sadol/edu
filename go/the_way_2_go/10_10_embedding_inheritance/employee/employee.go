package employee

import (
    "fmt"
    "../person"
)

type Employee struct {
	person.Person  "embedded anonymous → acts as an ancestor"
	Salary int "public field"
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Employee's id: %d.\n", e.Id())
	fmt.Printf("Employee's name: %s.\n", e.FirstName)
	fmt.Printf("Employee's surname: %s.\n", e.LastName)
	fmt.Printf("Employee's salary: %d.\n", e.Salary)
}

// pseudo constructor, order of creation is crucial here
func NewEmployee(id int, name string, surname string, salary int) (e *Employee) {
	e = new(Employee)
	e.Person = *person.NewPerson(id, name, surname)
	e.Salary = salary
	return
}
