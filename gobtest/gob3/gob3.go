/**
 * gob3.go is a practice for reading '***.gob' file
 */
package main

import (
	"encoding/gob"
	"fmt"
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

func main() {

	file, err := os.Open("D:\\Temp\\vcard.gob")
	if err != nil {
		fmt.Println("Reading file failed.")
	}

	var vcard VCard

	dec := gob.NewDecoder(file)
	errDec := dec.Decode(&vcard)

	if errDec != nil {
		fmt.Println("Decoding file failed.")
	}

	fmt.Printf("VCard is %v: \n", vcard)
}
