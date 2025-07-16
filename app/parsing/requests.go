package parsing

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Request struct {
	Method  string
	Target  string
	Version string
	Headers map[string]string
	Body    string
}

var ParseRequestError = errors.New("failed to parse HTTP Request")

func ParseRequest(input string) (Request, error) {
	splitInput := strings.SplitN(input, CLRF, 2)

	firstLine := strings.Split(splitInput[0], " ")
	if len(firstLine) != 3 {
		fmt.Println("Parse Request Error: first line not len of 3: ", firstLine)
		return Request{}, ParseRequestError
	}

	method := firstLine[0]
	target := firstLine[1]
	version := firstLine[2]

	headersAndBody := strings.SplitN(splitInput[1], CLRF+CLRF, 2)
	if len(headersAndBody) != 2 {
		fmt.Println("Parse Request Error: header + body not len of 2: ", headersAndBody)
		return Request{}, ParseRequestError
	}

	header := headersAndBody[0]
	body := headersAndBody[1]

	headerMap := make(map[string]string)

	if header != "" {
		headerSlice := strings.Split(header, CLRF)

		for i := 0; i < len(headerSlice); i++ {
			splitCurrentHeader := strings.Split(headerSlice[i], ": ")
			if len(splitCurrentHeader) != 2 {
				continue
			}

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

	logRequest(req)
	return req, nil
}

func logRequest(request Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] "+
		"%s %s %s\n", now, request.Method, request.Target, request.Version)
}
