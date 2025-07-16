package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

const CLRF = "\r\n"

type Response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
	Headers      map[string]string
	Body         string
}

func (receiver *Response) ToBytes() []byte {
	headerStrings := make([]string, 1)

	length := len(receiver.Body)
	receiver.Headers["Content-Length"] = strconv.Itoa(length)

	for k, v := range receiver.Headers {
		current := fmt.Sprintf("%s: %s%s", k, v, CLRF)
		headerStrings = append(headerStrings, current)
	}

	headerSection := strings.Join(headerStrings, "")

	res := fmt.Sprintf("%s %s %s%s%s\n%s",
		receiver.Version,
		receiver.StatusCode,
		receiver.ReasonPhrase,
		CLRF,
		headerSection,
		receiver.Body,
	)
	return []byte(res)
}

func GetNotFound(body string) Response {
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "404",
		ReasonPhrase: "Not Found",
		Body:         body,
		Headers:      make(map[string]string),
	}
}

func GetOk(body string) Response {
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "200",
		ReasonPhrase: "OK",
		Body:         body,
		Headers:      make(map[string]string),
	}
}

func GetBadRequest(body string) Response {
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "400",
		ReasonPhrase: "Bad Request",
		Body:         body,
		Headers:      make(map[string]string),
	}
}
