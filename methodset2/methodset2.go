package main

import "fmt"

// List is a slice of int
type List []int

// Appender is an interface
type Appender interface {
	Append(int)
}

// Append is a method of List which implements Appender interface
func (l *List) Append(val int) {
	*l = append(*l, val)
}

// CountInfo is a function which has an Appender interface parameter
func CountInfo(a Appender, start, end int) {

	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

// Lener is an interface
type Lener interface {
	Len() int
}

// Len is a method of List which implements Lener interface
func (l List) Len() int {
	return len(l)
}

// LongEnough is a function which has a Lener intreface parameter
func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {

	var lst List

	// 因为List的Len方法实现了Lener接口的Len方法，所以List的变量lst可以赋值给接口Lener的变量
	// 如果Lener接口有多个方法，List类型的方法必须实现Lener接口的所有方法才能将List的变量赋
	// 值给接口Lener的变量，否则会产生编译错误
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)

	// 由于List的Append方法实现了Appender接口发Append方法，所以List的变量plst可以赋值给接口Appender的变量
	CountInfo(plst, 1, 10)

	if LongEnough(plst) {

		fmt.Printf("- plst is long enough\n")
	}
}
