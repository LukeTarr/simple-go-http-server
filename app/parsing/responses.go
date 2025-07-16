package parsing

import (
	"fmt"
	"strconv"
	"strings"
)

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
	b := "<h1>Not Found</h1>"
	if body != "" {
		b = body
	}
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "404",
		ReasonPhrase: "Not Found",
		Body:         b,
		Headers:      map[string]string{"Content-Type": "text/html"},
	}
}

func GetOk(body string) Response {
	b := "<h1>Hello World</h1>"
	if body != "" {
		b = body
	}
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "200",
		ReasonPhrase: "OK",
		Body:         b,
		Headers:      map[string]string{"Content-Type": "text/html"},
	}
}

func GetBadRequest(body string) Response {
	b := "<h1>Bad Request</h1>"
	if body != "" {
		b = body
	}

	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "400",
		ReasonPhrase: "Bad Request",
		Body:         b,
		Headers:      map[string]string{"Content-Type": "text/html"},
	}
}
