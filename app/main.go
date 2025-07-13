package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/connections"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go connections.HandleTcpConnection(&conn)
	}

}
