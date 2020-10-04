package main

import (
	"fmt"
)

/**
* 该程序会产生死锁。
* ch <- 2执行时，相同通道ch的接收端还没有准备好，所以会一直阻塞不会执行下一条命令造成死锁。
* 此处如果将ch <- 2修改为go func() {ch <- 2}，并且在调用f1函数后名增加休眠命令使发送和接收端同步，则可避免死锁。
 */
func main() {

	ch := make(chan int)
	ch <- 2
	fmt.Println("sent data.")
	go f1(ch)
}

func f1(ch chan int) {
	fmt.Println(<-ch)
}
