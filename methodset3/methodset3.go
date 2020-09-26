package main

import "fmt"

// Ts is struct for test
type Ts struct {
	test int
}

// Ti is an interface for test
type Ti interface {
	tm()

	// 如果接口中新增一个方法tm2，且Ts没有实现该方法，则Ts的变量无法赋值给Ti接口的变量
	// tm2()
}

func (ts *Ts) tm() {
	fmt.Printf("This is a method of Ts which implements Ti interface's tm method.\n")
	fmt.Printf("The value of struct's variable test is: %d\n", ts.test)
}

func mt(ti Ti) {
	ti.tm()
}

func main() {

	t := &Ts{
		100,
	}

	mt(t)
}
