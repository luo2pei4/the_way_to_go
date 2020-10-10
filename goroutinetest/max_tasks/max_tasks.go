package main

import "fmt"

// MAXREQS 同时最大请求数
const MAXREQS = 50

var sem = make(chan int, MAXREQS)

// Request is a struct
type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {

	fmt.Println(r.a)
}

// handle是以协程方式执行的函数
// 每次以协程方式运行handle函数时会向sem通道中放入一个数
// 当sem通道中的数据量达到MAXREQS时，sem <- 1处理将会被阻塞，只有当其他协程中执行了<-sem后，才会将1放入sem通道，并执行process函数
func handle(r *Request) {
	sem <- 1
	process(r)
	<-sem
}

func server(requestChan chan *Request) {

	for {
		request := <-requestChan
		go handle(request)
	}
}

func main() {
	service := make(chan *Request)
	go server(service)
}
