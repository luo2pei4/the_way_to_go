package main

import "fmt"

func a() {
	fmt.Println("enter func a")
	defer fmt.Println("leave func a")
	panic("panic in func a")
}

func b() {
	fmt.Println("enter func b")
	defer fmt.Println("leave func b")
	a()
}

func c() {
	fmt.Println("enter func c")
	defer fmt.Println("leave func c")
	b()
}

func d() {
	fmt.Println("enter func d")
	defer fmt.Println("leave func d")
	c()
}

func protect(g func()) {
	defer func() {
		fmt.Println("Done")
		if err := recover(); err != nil {
			fmt.Printf("run time panic: %v", err)
		}
	}()
	fmt.Println("Start")
	g()
}

func main() {

	protect(d)
}
