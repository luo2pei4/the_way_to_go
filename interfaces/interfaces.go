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

// Area is a method of Square which implements Shaper interface
func (s *Square) Area() float32 {
	return s.side * s.side
}

func main() {

	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1

	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
