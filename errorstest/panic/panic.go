package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Starting the program")
	// panic("A server error occurred: stopping the program!")
	// fmt.Println("Ending the program") // unreachable code

	user := os.Getenv("path")
	if user == "" {
		panic("Unknown user: no value for $USER")
	}

	fmt.Printf("%s", user)
}
