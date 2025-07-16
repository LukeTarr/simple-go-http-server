package parsing

import (
	"strings"
)

type Request struct {
	Method  string
	Target  string
	Version string
	Headers map[string]string
	Body    string
}

type ParseRequestError struct{}

func (receiver ParseRequestError) Error() string {
	return "Failed to parse HTTP Request"
}

func ParseRequest(input string) (Request, error) {
	splitInput := strings.SplitN(input, "\r\n", 2)

	firstLine := strings.Split(splitInput[0], " ")
	method := ""
	target := ""
	version := ""

	if len(firstLine) == 3 {
		method = firstLine[0]
		target = firstLine[1]
		version = firstLine[2]
	} else {
		return Request{}, ParseRequestError{}
	}

	headersAndBody := strings.SplitN(splitInput[1], "\r\n\r\n", 2)
	header := headersAndBody[0]
	body := headersAndBody[1]

	headerMap := make(map[string]string)

	if header != "" {
		headerSlice := strings.Split(header, "\r\n")

		for i := 0; i < len(headerSlice); i++ {
			splitCurrentHeader := strings.Split(headerSlice[i], ": ")
			key := splitCurrentHeader[0]
			value := splitCurrentHeader[1]

			headerMap[key] = value
		}
	}

	req := Request{
		Method:  method,
		Target:  target,
		Version: version,
		Headers: headerMap,
		Body:    body,
	}
	return req, nil
}
