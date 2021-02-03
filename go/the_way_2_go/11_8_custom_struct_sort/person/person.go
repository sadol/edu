/*
`sort.Sorter' interface for `Persons' is implicit provided that `Persons'
is equiped with correct set of methods.
*/
package person

const DEFAULTSIZE = 5

type Person struct {
	firstName string
	lastName  string
}

func NewPerson(name, surname string) (retPerson *Person) {
	retPerson = new(Person)
	retPerson.firstName = name
	retPerson.lastName = surname
	return
}

type Persons struct {
	data []Person
}

func (ps *Persons) Len() int {
	return len(ps.data)
}

func (ps *Persons) Less(index1, index2 int) bool {
	return ((ps.data[index1].lastName + ps.data[index1].firstName) <
		(ps.data[index2].lastName + ps.data[index2].firstName))
}

func (ps *Persons) Swap(index1, index2 int) {
	ps.data[index1], ps.data[index2] = ps.data[index2], ps.data[index1]
}

// pretty printing ...
func (ps *Persons) List() string {
    strToReturn := "["
	for index, value := range ps.data {
		strToReturn += "\"" + value.firstName + " " + value.lastName + "\""
		if index < ps.Len()-1 {
			strToReturn += " "
		}
	}
	strToReturn += "]"
	return strToReturn
}

// stringificator
func (ps *Persons) String() string {
	return "Persons → " + ps.List()
}

// inserts object into a list of people
func (ps *Persons) InsertPerson(personToAdd *Person, position int) (ret bool) {
	if position >= ps.Len() {
		ret = false
	} else {
        ps.data[position] = *personToAdd
        ret = true
    }
    return
}

// factory variadic function
func NewPersons(size ...int) (retPersons *Persons) {
	retPersons  = new(Persons)
    numberOfArgs := len(size)
    lenOfArray := 0

	if numberOfArgs == 0 {
		lenOfArray = DEFAULTSIZE
	} else { // set of user supplied values (may be only one element
		for _, value := range size {
			lenOfArray += value
		}
	}

	retPersons.data = make([]Person, lenOfArray)
	return
}
