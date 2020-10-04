package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 1)

	go func() {

		time.Sleep(1e9)
		ch <- 1
	}()

	select {
	case <-time.After(5e8):
		fmt.Println("Timeout.")
		break
	case <-ch:
		fmt.Println("hello world.")
	}

	fmt.Println("End of this program")
}
