// employee "package" with one method
package employee

type Employee struct {
	Salary float32
}

// reciever (in first parents) is a pointer to the type (in this case type is a
// struct)
func (emp *Employee) GiveRaise(howMuchPercent float32) {
	emp.Salary = ((howMuchPercent / 100) * emp.Salary) + emp.Salary
}
