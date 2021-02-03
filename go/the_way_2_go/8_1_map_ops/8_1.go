// initial fun with golang maps

package main

import "fmt"

func main() {
    // maps are dynamic but you can prepare initial size of a map to avoid
    // unnecessary calulations; capacity (2nd arg of `make',initial size)
    // can be 7 in this case
    var	days = make(map[string]int, 7)
    days["monday"] = 1
    days["tuesday"]= 2
    days["thursday"] = 3
    days["wednesday"] = 4
    days["friday"] = 5
    days["saturday"] = 6
    days["sunday"] = 7

    /*
    or:
    days := map[string]int{"monday": 1, "tuesday": 2, ... }
    */

    // print keys & vals
	for key, value := range days {
		fmt.Printf("%s is %d day of the week (in EUROPE!).\n", key, value)
	}

    // print keys only
	for key := range days {
		fmt.Printf("`%s' key is present in this map.\n", key)
	}

    // check if map index is present; golang idiom
	if _, isPresent := days["tuesday"]; isPresent {
		fmt.Println("`thuesday' is present in the map.")
	} else {
		fmt.Println("`thuesday' is NOT present in the map.")
	}

    if _, isPresent := days["trumpfday"]; isPresent {
		fmt.Println("`trumpfday' is present in the map.")
	} else {
		fmt.Println("`trumpfday' is NOT present in the map.")
	}

    // removing a key from a map; attempt to remove non existent key is SILENT
    delete(days, "trumpfday")

    // is it weekend yet?
    if isWeekend(days["saturday"]) {
        fmt.Println("Today is weekend day;(")
    } else {
        fmt.Println("Today is not weeked yet:)")
    }
}

// weekend check
func isWeekend(dayID int) (weekend bool) {
    if dayID == 6 || dayID == 7 {
        weekend = true
    } else {
        weekend = false
    }
    return
}
