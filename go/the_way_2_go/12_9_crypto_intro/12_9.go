package main

import (
	"crypto/sha512"           // do not use md5, sha1 and other broken hashgens
	"fmt"
)

func main() {
	hasher := sha512.New512_256()                      // step 1. create hasher
	textToHash := []byte("ala ma kota")         // step 2. prepare text to hash
	fmt.Printf("Text to hash → %v.\n", textToHash)
	checksum := sha512.Sum512_256(textToHash)         // step 3. print checksum
	fmt.Printf("Checksum → %v.\n", checksum)
	hasher.Reset()                           // step 4. reseting hasher buffers
}
