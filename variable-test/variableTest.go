package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

// 注：
// 1.值类型的变量保存在栈中
// 2.被引用的变量保存在堆中
