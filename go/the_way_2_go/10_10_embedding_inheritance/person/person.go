package person

import "../base"

type Person struct {
	base.Base   "embedded anonymous → acts as an ancestor"
	FirstName string "public field"
	LastName  string "public field"
}

// `constructor'
func NewPerson (i int, firstName string, lastName string) (p *Person) {
    // using base struct setter
    p = new(Person)
    p.Base = *base.NewBase(i)
    p.FirstName = firstName
    p.LastName = lastName
    return
}
