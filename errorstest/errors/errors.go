package main

import (
	"errors"
	"fmt"
)

var errorNotFound error = errors.New("Not found error")

func main() {
	fmt.Println(errorNotFound)
}
