package main

import (
	"fmt"
	"github.com/LukeTarr/simple-go-http-server/app/connections"
	"net"
	"os"
)

func main() {
	fmt.Println("*** Server Starting ***")
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("*** Server successfully listening on 0.0.0.0:4221 ***")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("*** Accepting new connection ***")
		go connections.HandleTcpConnection(&conn)
	}

}
