package connections

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/LukeTarr/simple-go-http-server/app/parsing"

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
		fmt.Println("Error Reading, ", err.Error())
	}

	resp := parsing.GetOk("<h1>Hello World!</h1>")

	requestBytes := string(bytes.TrimRight(input, "\x00"))
	req, err := parsing.ParseRequest(requestBytes)
	if errors.Is(err, parsing.ParseRequestError{}) {
		resp = parsing.GetBadRequest("<h1>Bad Request</h1>")
	}

	resp.Headers = map[string]string{"Content-Type": "text/html"}

	if req.Target != "/" {
		resp = parsing.GetNotFound("<h1>Not Found</h1>")
	}

	_, err = conn.Write(resp.ToBytes())
	if err != nil {
		fmt.Println("Error Writing, ", err.Error())
	}

}
