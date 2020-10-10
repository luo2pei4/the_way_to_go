package main

import "fmt"

// Request is a struct
type Request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

// 实际进行请求处理的函数
func run(op binOp, req *Request) {

	// 将请求中的数据传给外部传入的binOp类型的函数
	// 将binOp函数执行后的返回值，发送到请求对象自身的通道中
	// 在对应的接收处理执行前，函数将在这个位置阻塞
	req.replyc <- op(req.a, req.b)
}

// 函数的参数是binOp函数和请求(Request结构体)的通道
func server(op binOp, reqChan chan *Request) {

	// 无线循环用于接收请求通道的中的请求，并通过协程的方式处理传入的请求
	for {

		// 从请求通道中接收请求数据(Request对象的指针)
		// **当请求通道中没有数据时，程序将在这个位置阻塞**
		req := <-reqChan

		// 将传入的函数和请求数据传给run函数，并通过协程的方式执行run函数
		go run(op, req)
	}
}

// 启动服务
// 创建了一个请求(Request结构体)的通道
// 启动一个服务的协程，并将binOp函数和创建的请求通道传给server函数
// 返回请求的通道
func startService(op binOp) chan *Request {
	reqChan := make(chan *Request)
	go server(op, reqChan)
	return reqChan
}

func main() {

	// 调用startService函数，传入一个函数，并返回一个请求通道对象
	reqChan := startService(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]Request

	// 循环创建请求对象，并将请求对象发送到请求通道中
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)

		// 当把请求对象放入请求通道时，server函数将会马上将这个请求对象取出，并通过协程方式执行run函数
		// 在run函数中将传入的函数的执行结果放入请求对象自己的通道中
		reqChan <- req
	}

	for i := N - 1; i >= 0; i-- {

		// 此处从请求消息对象自己的通道中获取数据并进行判断
		// 从请求消息对象自己的通道中获取数据时，run函数才真正的把计算结果放入这个通道，并结束run函数，协程处理结束。
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}

	fmt.Println("Done")
}
