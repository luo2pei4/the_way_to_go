package main

import "fmt"

// Shaper is an interface
type Shaper interface {
	Area() float32
}

// Square is a struct
type Square struct {
	side float32
}

// Rectangle is a struct
type Rectangle struct {
	length, width float32
}

// Area is a method of Square which implements Shaper interface
func (a *Square) Area() float32 {
	return a.side * a.side
}

// Area is a method of Rectangle which implements Shaper interface
func (a *Rectangle) Area() float32 {
	return a.length * a.width
}

func main() {

	a1 := new(Rectangle)
	a1.length = 6
	a1.width = 8
	a2 := &Square{5}

	shapes := []Shaper{a1, a2}

	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
