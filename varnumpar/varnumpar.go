package main

import "fmt"

func main() {

	x := min(5, 2, 1, 2, 4, 7, 8)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 3, 7, 2, 9}

	// 传递一个切片给变参函数时，采用[切片名称...]的格式
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)
}

func min(s ...int) int {

	if len(s) == 0 {

		return 0
	}

	min := s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}
