package main

import (
	"fmt"
	"time"
)

func main() {

	defer executionTime(time.Now().UnixNano())
	fmt.Println("this is a test")
	time.Sleep(1000000)
}

func executionTime(start int64) {
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
