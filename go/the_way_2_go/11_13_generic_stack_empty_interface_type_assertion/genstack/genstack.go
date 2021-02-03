package genstack // generic stack

import (
	"errors"
	"strconv"
)

type Element interface{}
type Vector []Element

type GenStack interface {
	Len() int
	IsEmpty() bool
	Push(Element)
	Pop() (Element, error) //remove from stack and return
	Top() (Element, error) //return top element only (do not remove from the stack)
	String() string
}

type Stack struct {
	GenStack "(unnecessary) explicitly embedded generic stack interface"
	data     Vector "array of different type elements"
}

//------------Stack methods------------------
func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Top() (element Element, ok error) {
	if s.IsEmpty() {
		element, ok = nil, errors.New("Empty stack.")
	} else {
		element, ok = s.data[s.Len()-1], nil
	}
	return
}

//puts new element on the top of the stack
func (s *Stack) Push(element Element) {
	s.data = append(s.data, element)
}

//removes first element from the top of the stack
func (s *Stack) Pop() (element Element, ok error) {
	if element, ok = s.Top(); ok == nil {
		s.data = s.data[:s.Len()-1] // i hope garbage collector can clean this
	}
	return
}

//factory of the new stack (for expandig in the future)
func NewStack() *Stack {
	st := new(Stack)
	st.data = make(Vector, 0)
	return st
}

//stack's stringification function
func (s *Stack) String() string {
	var retString = "Stack(size=" + strconv.Itoa(s.Len()) + ") →\n"
	var strElement = ""
	for index, value := range s.data {
		// check type of the value
		switch value.(type) {
		case int:
			strElement = strconv.Itoa(value.(int))
		case float64:
			strElement = strconv.FormatFloat(value.(float64), 'f', 2, 64)
		case string:
			strElement = value.(string)
		default:
			strElement = "Stringer not implemented!!!"
		}
		retString += "\t" + strconv.Itoa(index) + ": " + strElement + "\n"
	}
	return retString
}
