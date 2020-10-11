package main

import (
	"flag"
	"fmt"
	"net"
)

const maxRead = 25

func checkError(err error, info string) {

	if err != nil {

		panic("ERROR: " + info + " " + err.Error())
	}
}

func initServer(hostAndPort string) *net.TCPListener {

	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	fmt.Println("Listening to: " + listener.Addr().String())
	return listener
}

func sayHello(conn net.Conn) {

	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := conn.Write(obuf)
	checkError(err, "Write: write "+string(wrote)+" bytes.")
}

func handlMsg(length int, err error, msg []byte) {

	if length > 0 {
		print("<", length, ":")

		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}

func connectionHandler(conn net.Conn) {

	connFrom := conn.RemoteAddr().String()
	fmt.Println("Connection from: ", connFrom)

	sayHello(conn)

	for {

		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0

		switch err {
		case nil:
			handlMsg(length, err, ibuf)
		// case os.EAGAIN:
		// 	continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}
