package server

import (
	"fmt"
	"github.com/LukeTarr/simple-go-http-server/app/parsing"
	"net"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	Port       int
	Address    string
	handlerMap []FunctionPath
}

func (receiver *Server) Listen() {
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
		go HandleTcpConnection(&conn, receiver.handlerMap)
	}

}

func (receiver *Server) Get(path string, handlerFunction func(request parsing.Request) string) {
	receiver.handlerMap = append(receiver.handlerMap, FunctionPath{
		Path:            strings.ToLower(path),
		Method:          "GET",
		HandlerFunction: handlerFunction,
	})
}

func (receiver *Server) Post(path string, handlerFunction func(request parsing.Request) string) {
	receiver.handlerMap = append(receiver.handlerMap, FunctionPath{
		Path:            strings.ToLower(path),
		Method:          "POST",
		HandlerFunction: handlerFunction,
	})
}

func (receiver *Server) Put(path string, handlerFunction func(request parsing.Request) string) {
	receiver.handlerMap = append(receiver.handlerMap, FunctionPath{
		Path:            strings.ToLower(path),
		Method:          "PUT",
		HandlerFunction: handlerFunction,
	})
}

func (receiver *Server) Patch(path string, handlerFunction func(request parsing.Request) string) {
	receiver.handlerMap = append(receiver.handlerMap, FunctionPath{
		Path:            strings.ToLower(path),
		Method:          "PATCH",
		HandlerFunction: handlerFunction,
	})
}

func (receiver *Server) Delete(path string, handlerFunction func(request parsing.Request) string) {
	receiver.handlerMap = append(receiver.handlerMap, FunctionPath{
		Path:            strings.ToLower(path),
		Method:          "DELETE",
		HandlerFunction: handlerFunction,
	})
}
