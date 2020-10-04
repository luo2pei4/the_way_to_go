package main

import (
	"fmt"
	"time"
)

// 在每个发送处理后增加了3秒的休眠，可以观察到接收者getData会打印出通道中的数据，
// 因此证明通道两端的处理是同步操作。
func sendData(ch chan string) {

	ch <- "Washington"
	time.Sleep(3 * 1e9)
	ch <- "Tripoli"
	time.Sleep(3 * 1e9)
	ch <- "London"
	time.Sleep(3 * 1e9)
	ch <- "Beijing"
	time.Sleep(3 * 1e9)
	ch <- "Tokyo"
	time.Sleep(3 * 1e9)

	close(ch)
}

func getData(ch chan string) {

	// for {
	// 	input, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(input)
	// }

	// 使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭
	for input := range ch {

		fmt.Println(input)
	}
}

func main() {

	ch := make(chan string)

	go sendData(ch)

	// 由于getData不是协程，所以只有当判断通道被关闭后才会主动退出，并结束main函数。
	// 如果用协程方式调用getData函数，程序会马上退出，不会打印任何数据。
	getData(ch)
}
