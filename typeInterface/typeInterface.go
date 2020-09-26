package main

import (
	"fmt"
	"math"
)

// Square is a struct
type Square struct {
	side float32
}

// Circle is a struct
type Circle struct {
	radius float32
}

// Shaper is a interface
type Shaper interface {
	Area() float32
}

// Area is a method of Square which implements Shaper interface
func (s *Square) Area() float32 {
	return s.side * s.side
}

// Area is a method of Circle which implements Shaper interface
func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {

	var areaIntf Shaper
	sq1 := &Square{5}

	areaIntf = sq1

	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf dose not contain a variable of type Circle")
	}
}
