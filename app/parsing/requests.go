package parsing

import (
	"strings"
)

type Request struct {
	Method  string
	Target  string
	Version string
	Headers RequestHeaders
}

type RequestHeaders = map[string]string

func ParseRequest(input string) Request {
	splitInput := strings.SplitN(input, "\r\n", 2)

	firstLine := strings.Split(splitInput[0], " ")
	method := firstLine[0]
	target := firstLine[1]
	version := firstLine[2]

	req := Request{
		Method:  method,
		Target:  target,
		Version: version,
		Headers: nil,
	}
	return req
}
