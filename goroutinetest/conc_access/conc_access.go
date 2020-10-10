package main

import (
	"fmt"
	"strconv"
	"time"
)

// Person is a struct
type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

func (p *Person) backend() {

	for f := range p.chF {
		f()
		fmt.Println("After execute f()", p.salary)
	}
}

// SetSalary is function
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() {
		p.salary = sal
	}
}

// Salary is a funcation
func (p *Person) Salary() float64 {

	// fChan := make(chan float64)
	// p.chF <- func() {
	// 	fChan <- p.salary
	// }
	// return <-fChan
	return p.salary
}

func (p *Person) String() string {
	// return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.salary, 'f', 2, 64)
}

// NewPerson is a function
func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func main() {

	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed: ")

	// 如果不休眠就马上打印Person对象，Person对象salary属性可能还没有被修改，backend函数相较于main函数是异步的，所以执行传入的匿名函数在时间上可能会落后于后面的打印处理。
	time.Sleep(1e9)
	fmt.Println(bs)
}
