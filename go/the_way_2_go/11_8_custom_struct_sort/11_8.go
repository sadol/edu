package main

import (
	"./person"
	"sort"
	"fmt"
)

func main() {
	var ps1 = person.NewPersons()
	var girl1 = person.NewPerson("Ala", "Bzik")
	var girl2 = person.NewPerson("Ala", "Gzik")
	var girl3 = person.NewPerson("Ola", "Bzik")
	var girl4 = person.NewPerson("Ola", "Gzik")
	var girl5 = person.NewPerson("Zuza", "Muza")
	_ = ps1.InsertPerson(girl1, 0)
	_ = ps1.InsertPerson(girl2, 1)
	_ = ps1.InsertPerson(girl3, 2)
	_ = ps1.InsertPerson(girl4, 3)
	_ = ps1.InsertPerson(girl5, 4)
	fmt.Printf("%v.\n", ps1)
	if sort.IsSorted(ps1) == false {
		fmt.Println()
		fmt.Println("Sorted:")
		sort.Sort(ps1)
		fmt.Printf("%v.\n", ps1)
	}
}
