// Vcard assignment

package main

import (
	"fmt"
	"time"
)

func main() {
	type Address struct {
		Coutry       string
		State        string
		City         string
		Street       string
		StreetNo     int
		ApartamentNo int
	}

	type Photo struct {
		PhotoBytes [4096]byte
	}

	type VCard struct {
		Name      string
		Adresses  *[]*Address
		Photos    *[]*Photo
		BirthDate time.Time
	}

	kowalskiPhoto1 := Photo{PhotoBytes: [4096]byte{1, 2, 3}}
	kowalskiPhoto2 := Photo{PhotoBytes: [4096]byte{4, 5, 6}}
	kowalskiPhotos := []*Photo{&kowalskiPhoto1, &kowalskiPhoto2}
	kowalskiAddress1 := Address{"Czechia", "Moravia", "Prague", "Havla", 12, 2}
	kowalskiAddress2 := Address{"Czechia", "Moravia", "Prague", "Karola", 5, 0}
	kowalskiAddresses := []*Address{&kowalskiAddress1, &kowalskiAddress2}
	myLocation, _ := time.LoadLocation("UTC")
	kowalski := VCard{"Jan Kowalski", &kowalskiAddresses, &kowalskiPhotos,
		time.Date(1920, time.Month(3), 23, 0, 0, 0, 0, myLocation)}
	fmt.Println(kowalski)
}
