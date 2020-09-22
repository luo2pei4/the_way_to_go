package main

import "fmt"

// testStruct is a test struct
type testStruct struct {
	str string
}

func main() {

	m1 := new(testStruct)
	fmt.Printf("new struct address: %p\n", m1)
	fmt.Printf("m1 address: %v\n", &m1)

	m2 := &testStruct{
		"this is a test",
	}

	fmt.Printf("initialize1 struct address: %p\n", m2)
	fmt.Printf("m2 address: %v\n", &m2)

	m3 := testStruct{
		"this is a test",
	}

	fmt.Printf("initialize2 struct address: %p\n", &m3)
	fmt.Printf("m3 address: %v\n", &m3)
}
