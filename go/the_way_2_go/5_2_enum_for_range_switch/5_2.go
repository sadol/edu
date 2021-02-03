// golang switch usage

package main

import "fmt"

//---- golang enums---------
type Month int

const (
	_         = iota                                               // discard 0
	JAN Month = iota                                            // start from 1
	FEB                                         // no need for FEB Month = iota
	MAR
	APR
	MAY
	JUN
	JUL
	AUG
	SEP
	OCT
	NOV
	DEC
	XXX // intentional error
)

type Season int

const (
    WINTER Season = iota                  // this time use 0 as a `winter' code
    SPRING
    SUMMER
    AUTUMN
    YYY                                                   // error code storing
)
//-------------------------

func main() {
	var tests = []Month{JAN, MAY, FEB, XXX}
	var seasonGot Season
	var errorGot bool = false
	fmt.Println()
	fmt.Println("`Season' testing:")

    // index, value = range <collection> {...}
    // `_' means ignore
	for _, val := range tests {
		seasonGot, errorGot = getSeason(val)
		fmt.Printf("Month: %d\tseason: %s\t ok?: %t.\n", int(val),
			translateSeason(seasonGot), errorGot)
	}
}

/*
    It's neccesary to use enumaration type here, feeding function with base type
    of the said enumeration will trigger compilation error;
	`break' clause is explicit in golang `switch'
*/
func getSeason(month Month) (Season, bool) {
	var output Season
	var err bool = true              // should be specialized `error' type here
	switch month { // does NOT have to be int or const
	case DEC, JAN, FEB: // but here must have the same type as switch operand
		output = WINTER
	case MAR, APR, MAY:
		output = SPRING
	case JUN, JUL, AUG:
		output = SUMMER
	case SEP, OCT, NOV:
		output = AUTUMN
	default: // double check of itself
        output = YYY
		err = false
	}
	return output, err
}

// translates `Season' enum codes to strings
func translateSeason ( inputSeason Season ) string {
    var seasonName string
    switch inputSeason {
    case WINTER:
        seasonName = "winter"
    case SPRING:
        seasonName = "spring"
    case SUMMER:
        seasonName = "summer"
    case AUTUMN:
        seasonName = "autumn"
    default:
        seasonName = "error"
    }
    return seasonName
}
