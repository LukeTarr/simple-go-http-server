package server

import (
	"fmt"
	"github.com/LukeTarr/simple-go-http-server/app/connections"
	"net"
	"os"
	"strconv"
)

type Server struct {
	Port    int
	Address string
}

func (receiver Server) Listen() {
	fmt.Println("*** Server Starting ***")

	p := strconv.Itoa(receiver.Port)
	addrPort := fmt.Sprintf("%s:%s", receiver.Address, p)

	listener, err := net.Listen("tcp", addrPort)
	if err != nil {
		fmt.Println("Failed to bind to port ", receiver.Port)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("*** Server successfully listening on ", addrPort, " ***")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go connections.HandleTcpConnection(&conn)
	}

}
