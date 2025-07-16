package server

import (
	"bytes"
	"fmt"
	"github.com/LukeTarr/simple-go-http-server/app/parsing"
	"io"

	"net"
)

func HandleTcpConnection(c *net.Conn, functionPaths []FunctionPath) {
	conn := *c
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error Closing connection")
			return
		}
	}()

	input := make([]byte, 16384)
	_, err := conn.Read(input)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error Reading, ", err.Error())
		}

	}

	resp := parsing.GetOk("")

	requestBytes := string(bytes.TrimRight(input, "\x00"))
	req, err := parsing.ParseRequest(requestBytes)

	didFindFunction := false

	for i := 0; i < len(functionPaths); i++ {

		if functionPaths[i].Path == req.Target {
			if functionPaths[i].Method == req.Method {
				didFindFunction = true
				returnBody := functionPaths[i].HandlerFunction(req)
				resp.Body = returnBody
			}
		}
	}

	if !didFindFunction {
		resp = parsing.GetNotFound("")
	}

	_, err = conn.Write(resp.ToBytes())
	if err != nil {
		fmt.Println("Error Writing, ", err.Error())
	}

}
