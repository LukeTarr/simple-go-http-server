package response

import (
	"fmt"
	"github.com/LukeTarr/simple-go-http-server/app/parsing"
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

func NewResponse(version, statusCode, reasonPhrase, body string) *Response {
	return &Response{
		Version:      version,
		StatusCode:   statusCode,
		ReasonPhrase: reasonPhrase,
		Headers:      make(map[string]string),
		Body:         body,
	}
}

// SetHeader sets a header key-value pair.
func (receiver *Response) SetHeader(key, value string) {
	receiver.Headers[key] = value
}

// SetBody sets the response body.
func (receiver *Response) SetBody(body string) {
	receiver.Body = body
}

func (receiver *Response) SetStatusCode(statusCode int) {
	receiver.StatusCode = strconv.Itoa(statusCode)
}

func (receiver *Response) ToBytes() []byte {
	headerStrings := make([]string, 1)

	length := len(receiver.Body)
	receiver.Headers["Content-Length"] = strconv.Itoa(length)

	for k, v := range receiver.Headers {
		current := fmt.Sprintf("%s: %s%s", k, v, parsing.CLRF)
		headerStrings = append(headerStrings, current)
	}

	headerSection := strings.Join(headerStrings, "")

	res := fmt.Sprintf("%s %s %s%s%s\n%s",
		receiver.Version,
		receiver.StatusCode,
		receiver.ReasonPhrase,
		parsing.CLRF,
		headerSection,
		receiver.Body,
	)
	return []byte(res)
}
