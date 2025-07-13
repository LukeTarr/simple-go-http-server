package connections

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/parsing"
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

	for {
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error Reading, ", err.Error())
			break
		}

		response := parsing.Response{
			Version:      "HTTP/1.1",
			StatusCode:   "200",
			ReasonPhrase: "OK",
		}

		_, err = conn.Write(response.ToBytes())
		if err != nil {
			fmt.Println("Error Writing, ", err.Error())
			break
		}

	}

}
