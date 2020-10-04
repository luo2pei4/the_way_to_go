package main

import (
	"fmt"
	"time"
)

func main() {

	stream := pump()

	// 此处suck函数即使不是协程也会持续接收pump函数中压入通道的数据，能把程序搞挂
	// suck(stream)

	// 如果suck函数为协程，程序将会在输出257001时结束
	go suck(stream)
	time.Sleep(1e9)
}

func pump() chan int {

	ch := make(chan int)
	go func() {

		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func suck(ch chan int) {

	for {
		fmt.Println(<-ch)
	}
}
