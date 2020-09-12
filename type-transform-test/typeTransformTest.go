package main

import (
	"fmt"
	"math"
)

// IntFromFloat64 is a test
func IntFromFloat64(x float64) int {

	if math.MinInt64 <= x && x <= math.MaxInt64 {

		whole, fraction := math.Modf(x)

		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int64 range.", x))
}

func main() {

	println(IntFromFloat64(15.186))
}
