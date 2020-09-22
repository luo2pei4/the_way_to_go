package main

import (
	"fmt"
	"time"
)

// myTime is a struct, use to implement self method of time.Time
type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {

	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())
}
