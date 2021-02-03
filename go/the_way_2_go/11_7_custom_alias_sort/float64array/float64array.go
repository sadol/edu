package float64array

import (
	"math/rand"
	"strconv"
	"time"
)

// default size of the `data' field of the `Float64Array'
const DEFAULTSIZE = 25

/* this is the usual Sorter interface to use with the sort.Sort method;
   Sorter interface need not to be implemented explicitly as long as
   methods of this interface are implemented for certain struct.
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
*/

type Float64Array []float64

func (fa *Float64Array) Len() int {
	return len(*fa)
}

func (fa *Float64Array) Less(index1, index2 int) bool {
	return (*fa)[index1] < (*fa)[index2]
}

func (fa *Float64Array) Swap(index1, index2 int) {
	(*fa)[index1], (*fa)[index2] = (*fa)[index2], (*fa)[index1]
}

// float64array pretty printing
func (fa *Float64Array) List() (strToReturn string) {
	strToReturn = "["
	for index, value := range (*fa) {
		strToReturn += strconv.FormatFloat(value, 'f', 2, 64)
		if index < fa.Len() {
			strToReturn += " "
		}
	}
	strToReturn += "]"
	return
}

// stringificator
func (fa *Float64Array) String() string {
	return "Float64Array (size = " + strconv.Itoa(fa.Len()) + ") → " + fa.List()
}

// fills flaot64array with random values
func (fa *Float64Array) Fill() {
	rand.Seed(time.Now().UTC().UnixNano())
    fillVal := rand.NormFloat64()
	for index, _ := range (*fa) {
		(*fa)[index] = fillVal
		fillVal = rand.NormFloat64()
	}
}

//factory variadic function
func NewFloat64Array(size ...int) *Float64Array {
    numberOfArgs := len(size)
    lenOfArray := 0

	if numberOfArgs == 0 {
		lenOfArray = DEFAULTSIZE
	} else {
		for _, value := range size {
			lenOfArray += value
		}
	}

    newArray := Float64Array(make([]float64, lenOfArray))
	newArray.Fill()
	return &newArray
}
