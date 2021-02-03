// defining and using aliases

package main

import "fmt"

type Rope string // alias is a brand new type (with new interface if needed)

func main() {
	// Geene Hackman in the 'A Bridge Too Far' as Gen.Sosabowski
	var line Rope = "Ciągnij linę, sznur!?!"
	fmt.Printf("var line Rope = %s\n", line)
}
