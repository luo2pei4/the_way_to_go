package main

import "fmt"

func main() {

	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	fmt.Printf("Address of slTo before call copy function is: %p\n", slTo)
	n := copy(slTo, slFrom)
	fmt.Printf("Address of slTo after call copy function is: %p\n", slTo)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	fmt.Printf("Address of sl3 before call append function is: %p\n", sl3)
	sl3 = append(sl3, 4, 5, 6, 7)
	fmt.Printf("Address of sl3 before call append function is: %p\n", sl3)
	fmt.Println(sl3)
}
