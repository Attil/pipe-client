package main

import (
	"flag"

	"fmt"

	winio "github.com/Microsoft/go-winio"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "mock", "name for the pipe")

	address := "//./pipe/" + name

	conn, err := winio.DialPipe(address, nil)

	if err != nil {
		fmt.Println("Could not connect to the selected pipe")
		return
	}

	writeBuffer := []byte("{\"type\": nopenopenopenopenope}")
	_, err = conn.Write(writeBuffer)

	if err != nil {
		fmt.Println("Could not send the message")
		return
	}

	var readBuffer [4096]byte
	bytesRead, err := conn.Read(readBuffer[0:])

	if err != nil {
		fmt.Println("Could not receive the response")
		return
	}

	response := string(readBuffer[:bytesRead])

	fmt.Println("Got response:", response)

	return
}
