// crude and unsafe implementation of the golang stack

package arrstack

import "strconv"

// very crude stack implementation
type Arr4Stack struct {
	size  int   "size of the stack"
	top   int   "pointer to the top of the stack"
	stack []int "stack itself"
}

//contructor of the new stack
func NewStack(size int) *Arr4Stack {
	st := new(Arr4Stack)
	st.size = size
	st.top = 0
	for i := 0; i < size; i++ {
		st.stack = append(st.stack, 0)
	}
	return st
}

//removes first element from the top of the stack
func (a *Arr4Stack) Pop() (element int, ok bool) {
	switch {                                // nice GO idiom: switch on nothing
	case a.top <= 0 || a.top > a.size: //error of some kind OR empty stack (a.top = -1)
		element, ok = 0, false
	case a.top > 0 && a.top <= a.size:
		element, ok = a.stack[a.top-1], true
		a.stack[a.top-1] = 0                                    // clear memory
		a.top--
	}
	return
}

//puts new element on the top of the stack
func (a *Arr4Stack) Push(element int) (retValue bool) {
	switch {                                // nice GO idiom: switch on nothing
	case a.top == a.size:                                         // full stack
		retValue = false
	case a.top == -1:               // reuse old stack which was cleared before
		a.top = 0
		a.stack[0] = element
		retValue = true
	case a.top < -1:                              // some kind of a grave error
		retValue = false
	default:                                            // stack is operational
		a.stack[a.top] = element
		a.top++
		retValue = true
	}
	return
}

//stack's stringification function
func (s *Arr4Stack) String() string {
	return "Stack(size=" + strconv.Itoa(s.size) + ", top=" + strconv.Itoa(s.top) +
		") → [0:" + strconv.Itoa(s.stack[0]) + "] [1:" + strconv.Itoa(s.stack[1]) +
		"] [2:" + strconv.Itoa(s.stack[2]) + "] [3:" + strconv.Itoa(s.stack[3]) + "]"
}
