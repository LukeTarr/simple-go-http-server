package connections

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/LukeTarr/simple-go-http-server/app/parsing"
	"io"

	"net"
)

func HandleTcpConnection(c *net.Conn) {
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

	if errors.Is(err, parsing.ParseRequestError) {
		resp = parsing.GetBadRequest("")
	}

	if req.Target != "/" {
		resp = parsing.GetNotFound("")
	}

	_, err = conn.Write(resp.ToBytes())
	if err != nil {
		fmt.Println("Error Writing, ", err.Error())
	}

}
