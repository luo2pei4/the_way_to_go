package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(1e9)
	boom := time.After(5e9)

	for {

		select {
		case v := <-tick:
			fmt.Println("tick.", v)
		case v := <-boom:
			fmt.Println("boom.", v)
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e8)
		}
	}
}
