package main

import "fmt"

// TwoInts is a struct
type TwoInts struct {
	a, b int
}

func main() {

	two1 := TwoInts{1, 2}

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

// AddThem is a method of struct TwoInts
func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

// AddToParam is a method of struct TwoInts
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
