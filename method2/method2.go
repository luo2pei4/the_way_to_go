package main

import "fmt"

// IntVector is a slice of int vector
type IntVector []int

// Sum is a method of slice IntVector
func (v IntVector) Sum() (s int) {

	fmt.Printf("init return value s is %v\n", s)

	for _, x := range v {
		s += x
	}
	return
}

func main() {

	fmt.Println(IntVector{1, 2, 3}.Sum())
}
