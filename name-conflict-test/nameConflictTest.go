package main

import "fmt"

// AA is a struct
type AA struct{ a int }

// B is a struct
type B struct{ a, b int }

// C is a struct
type C struct {
	AA // 匿名字段
	B  // 匿名字段
}

func main() {

	a := AA{1}
	b := B{2, 3}
	c := C{a, b}

	// fmt.Println(c.a) 变量a在结构体AA和B中都存在，此处会引起二义性错误
	fmt.Println(c.AA.a)
	fmt.Println(c.b)
}
