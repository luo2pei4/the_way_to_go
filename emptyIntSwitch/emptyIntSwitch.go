package main

import "fmt"

type specialString string

var whatIsThis specialString = "hello"

// TypeSwitch is a function
func TypeSwitch() {

	// 匿名函数的参数为空指针，可以接受任何类型的值
	testFunc := func(any interface{}) {

		// 通过“变量.(type)的方式获取变量的类型”
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special string!", v)
		default:
			fmt.Println("unknow type!")
		}
	}
	testFunc(whatIsThis)
}

func main() {
	TypeSwitch()
}
