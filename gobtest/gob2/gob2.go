package main

import (
	"encoding/gob"
	"log"
	"os"
)

// Address is a struct
type Address struct {
	Type    string
	City    string
	Country string
}

// VCard is a struct
type VCard struct {
	FirstName string
	LastName  string
	Addresser []*Address
	Remark    string
}

var content string

func main() {

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belguim"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	file, _ := os.OpenFile("D:\\Temp\\vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoding gob")
	}
}
