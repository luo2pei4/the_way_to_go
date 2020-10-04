package main

import (
	"fmt"
	"time"
)

func pump1(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1e9)
	}
	// 此处关闭通道，如果select中没有做通道关闭判断会一直接收到0值
	close(ch)
}

func pump2(ch chan int) {

	for i := 10; i > 0; i-- {
		ch <- i
		time.Sleep(1e9)
	}
	close(ch)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)

	for {

		// 如果通道被关闭，此处会始终接收到0值或nil，此处对通道关闭做了判断，并做线程休眠
		select {
		case v, ok := <-ch1:

			if !ok {

				fmt.Println("channel 1 was closed.")
				time.Sleep(1e9)

			} else {

				fmt.Printf("Received on channel 1: %d\n", v)
			}

		case v, ok := <-ch2:

			if !ok {

				fmt.Println("channel 2 was closed.")
				time.Sleep(1e9)

			} else {

				fmt.Printf("Received on channel 2: %d\n", v)
			}

		default:
			fmt.Println("Waiting for data...")
			time.Sleep(1e9)
		}
	}
}
