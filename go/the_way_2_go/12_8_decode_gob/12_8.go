package main

import (
	"encoding/gob" // generic golang serialization (binary in transit)
	"fmt"
	"log"
	"os"
)

type AddressType int

const (
	HOME AddressType = iota
	WORK
)

type Address struct {
	Type    AddressType
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	var cardFile *os.File // must be here to defer close
	homeAddress := &Address{HOME, "Newer York", "Russia"}
	workAddress := &Address{WORK, "Sina Dupa", "Russia"}
	card := VCard{"Wowa", "Botox", []*Address{homeAddress, workAddress}, "toxic"}
	// using gob encoder to serialize into the stream
	if cardFile, err := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		log.Fatal("open file error:", err)
	} else {
	    defer cardFile.Close()
		encoder := gob.NewEncoder(cardFile)          // step 1. create new encoder
		if err := encoder.Encode(card); err != nil { // step 2. encode (into the file)
			log.Fatal("encode error:", err)
		} else {
			fmt.Println("GOB file created.")
		}
	}

	var outputCard VCard
	if cardFile, err := os.Open("vcard.gob"); err != nil {
		log.Fatal("open file error:", err)
	} else {
		decoder := gob.NewDecoder(cardFile)                          // step 3. create new decoder
		if err := decoder.Decode(&outputCard); err != nil {          // step 4. decode into the struct
			log.Fatal("decode error:", err)
		} else {
			fmt.Printf("Decoded value : %v.\n", outputCard)          // step 5. do sth with decoded struct
			fmt.Printf("Address1: %v, address2: %v.\n", *outputCard.Addresses[0],
				*outputCard.Addresses[1])
		}
	}
}
