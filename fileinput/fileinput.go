package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	inputFile, inputError := os.Open("D:\\Code\\other\\priorityRules.xml")

	if inputError != nil {
		fmt.Println("Open file failed.")
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {

		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("%v", inputString)

		if readerError == io.EOF {
			return
		}
	}
}
