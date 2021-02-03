// inserting slice into another slice at certian place

package main

import "fmt"

func main() {

	byteStringBase := make([]byte, 5, 10)
	byteStringBase[0] = '0'
	byteStringBase[1] = '1'
	byteStringBase[2] = '2'
	byteStringBase[3] = '3'
	byteStringBase[4] = '4'
	byteStringInsert := make([]byte, 3, 3)
	byteStringInsert[0] = 'a'
	byteStringInsert[1] = 'b'
	byteStringInsert[2] = 'c'
	fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert)
	where := 15
    fmt.Printf("Insert position: %d; expected failure: %t.\n", where, true)
	byteStringReturn, ok := InsertStringSlice(byteStringBase, byteStringInsert, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}

	fmt.Println()
    fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert)
	where = 15
    fmt.Printf("Insert position: %d; expected failure(2n version): %t.\n", where, false)
	byteStringReturn, ok = InsertStringSlice2(byteStringBase, byteStringInsert, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}

	fmt.Println()
	fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert)
	where = 5
    fmt.Printf("Insert position: %d; expected failure: %t.\n", where, true)
	byteStringReturn, ok = InsertStringSlice(byteStringBase, byteStringInsert, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}

    fmt.Println()
	fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert)
	where = 5
    fmt.Printf("Insert position: %d; expected failure(2nd version): %t.\n", where, false)
	byteStringReturn, ok = InsertStringSlice2(byteStringBase, byteStringInsert, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}

	fmt.Println()
	fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert)
	where = 2
    fmt.Printf("Insert position: %d; expected failure: %t.\n", where, false)
	byteStringReturn, ok = InsertStringSlice(byteStringBase, byteStringInsert, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}

	fmt.Println()
	fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert)
	where = 2
    fmt.Printf("Insert position: %d; expected failure(2nd version): %t.\n", where, false)
	byteStringReturn, ok = InsertStringSlice2(byteStringBase, byteStringInsert, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}

	fmt.Println()
	byteStringInsert1 := make([]byte, 10, 30)
	byteStringInsert1[0] = 'a'
	byteStringInsert1[1] = 'b'
	byteStringInsert1[2] = 'c'
	byteStringInsert1[3] = 'd'
	byteStringInsert1[4] = 'e'
	byteStringInsert1[5] = 'f'
	byteStringInsert1[6] = 'g'
	byteStringInsert1[7] = 'h'
	byteStringInsert1[8] = 'i'
	byteStringInsert1[9] = 'j'
	fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert1)
	where = 3
    fmt.Printf("Insert position: %d; expected failure: %t.\n", where, false)
	byteStringReturn, ok = InsertStringSlice(byteStringBase, byteStringInsert1, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n ", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}

	fmt.Println()
    fmt.Printf("Base byte string: %q.\n ", byteStringBase)
	fmt.Printf("Insert byte string: %q.\n ", byteStringInsert1)
	where = 3
    fmt.Printf("Insert position: %d; expected failure(2nd version): %t.\n", where, false)
	byteStringReturn, ok = InsertStringSlice2(byteStringBase, byteStringInsert1, where)
	if ok {
		fmt.Printf("Modified byte string (where=%d): %q.\n ", where, byteStringReturn)
	} else {
		fmt.Println("Error ocurred.")
	}
}

/*
Inserts slice into another slice.
There is possibility that len != cap for origSlice or insertSlice respectively,
but function ignores this completely.
insertStart is AN INDEX (not a position in an array).
*/
func InsertStringSlice (origSlice []byte, insertSlice []byte, insertStart int) (returnSlice []byte, ok bool) {
	insertEnd := insertStart + len(insertSlice)
	returnSlice = make([]byte, len(origSlice) + len(insertSlice))

	if insertStart > len(origSlice) || insertStart < 0 {
		ok = false // error
	} else {
        copy(returnSlice[:insertStart], origSlice[:insertStart])
		copy(returnSlice[insertStart:insertEnd], insertSlice)
        copy(returnSlice[insertEnd:], origSlice[insertStart:])
        ok = true
	}
    return
}

/*
Second version of the function above, using `cap' instead of `len'.
Caller is reponsible for avoiding overflow errors using proper args of this
function.
*/
func InsertStringSlice2 (origSlice []byte, insertSlice []byte, insertStart int) (returnSlice []byte, ok bool) {
	if insertStart < 0 {
		ok = false // error
	} else {
        // create array here , after initial insertStart testing
        insertEnd := insertStart + cap(insertSlice)
        sizeOfReturnSlice := 0
        if (insertStart <= cap(origSlice)) { // normal case
            sizeOfReturnSlice = cap(origSlice) + cap(insertSlice)
            returnSlice = make([]byte, sizeOfReturnSlice)
            copy(returnSlice[:insertStart], origSlice[:insertStart])
            copy(returnSlice[insertStart:insertEnd], insertSlice)
            copy(returnSlice[insertEnd:], origSlice[insertStart:])
        } else { // extreme case
            sizeOfReturnSlice = cap(insertSlice) + insertStart
            returnSlice = make([]byte, sizeOfReturnSlice)
            copy(returnSlice[:insertStart], origSlice)
            copy(returnSlice[insertStart:insertEnd], insertSlice)
        }
        ok = true
	}
    return
}
