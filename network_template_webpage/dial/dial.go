package main

import (
	"fmt"
	"net"
)

func checkConnection(conntype string, conn net.Conn, err error) {

	fmt.Printf("connection type is: %v\n", conntype)

	if err != nil {

		fmt.Printf("error %v connecting!\n", err)

	} else {

		fmt.Printf("Connectiong is made with %v\n", conn)
	}
}

func main() {

	conn, err := net.Dial("tcp", "localhost:80")
	checkConnection("tcp", conn, err)

	conn, err = net.Dial("udp", "localhost:80")
	checkConnection("udp", conn, err)
}
