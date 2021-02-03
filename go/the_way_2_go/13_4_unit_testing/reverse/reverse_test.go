package reverse // the same package

import (
	"fmt"
	"strings"
	"testing" // obligatory !!!
)

func TestReverse(t *testing.T) { // name should be TestXxx and function header has to contain *testing.T address
	argument1 := "Ala"
	shouldReturn1 := "alA"
	if strings.Compare(Reverse(argument1), shouldReturn1) != 0 {
		t.Log(fmt.Sprintf("`%v' is NOT reversed `%v'!", shouldReturn1, argument1))
		t.Fail()
	}
}
