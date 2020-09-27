package main

import "fmt"

// Tester is an interface
type Tester interface {
	test()
}

// Ts is a struct
type Ts struct {
	testString string
}

func (ts *Ts) test() {
	fmt.Println(ts.testString)
}

func main() {

	var tester Tester
	ts := &Ts{"This is a test."}
	tester = ts

	if sv, ok := tester.(Tester); ok {

		sv.test()
	}
}
