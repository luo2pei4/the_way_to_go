package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	var (
		host   = "www.apache.org"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 4096)
		read   = true
		count  = 0
	)

	conn, err := net.Dial("tcp", remote)

	if err != nil {

		fmt.Println("Dialing failed,", err.Error())
		return
	}

	// conn对象实现了io.Writer接口的write方法
	if sv, ok := conn.(io.Writer); ok {

		fmt.Println("conn implements io.Writer interface.")
		fmt.Println(sv)
	}

	// conn对象实现了net.Conn接口的write方法
	if sv, ok := conn.(net.Conn); ok {

		fmt.Println("conn implements net.Conn interface.")
		fmt.Println(sv)
	}

	// 在WriteString方法中调用了io.Writer接口的write方法，如果不是write方法，则无法将conn对象作为参数传递给io.WriteString
	io.WriteString(conn, msg)

	for read {

		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Println(string(data[0:count]))
	}

	conn.Close()
}
