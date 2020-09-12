package main

import "fmt"

// init 函数是一个特殊的函数，每个含有该函数的包都会首先执行这个函数。
func init() {

	fmt.Println("Execute init function.")
}

func main() {

	fmt.Println("Execute main function.")
}
