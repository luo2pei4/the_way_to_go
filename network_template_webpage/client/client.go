package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:50000")

	if err != nil {

		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	tirmmedClient := strings.Trim(clientName, "\r\n")

	for {

		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		tirmmedInput := strings.Trim(input, "\r\n")

		if tirmmedInput == "Q" {
			return
		}

		temp := string(tirmmedClient + " says: " + tirmmedInput)

		_, err := conn.Write([]byte(temp))

		if err != nil {

			fmt.Println("Send message \"" + temp + "\" failed. Error detail is: \n" + err.Error())
		}
	}
}
