package parsing

import "fmt"

const CLRF = "\r\n"

type Response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
	Body         string
}

func (receiver Response) ToBytes() []byte {
	res := fmt.Sprintf("%s %s %s%s%s\n%s",
		receiver.Version,
		receiver.StatusCode,
		receiver.ReasonPhrase,
		CLRF,
		CLRF,
		receiver.Body,
	)
	return []byte(res)
}

func GetNotFound() Response {
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "404",
		ReasonPhrase: "Not Found",
	}
}

func GetOk(body string) Response {
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "200",
		ReasonPhrase: "OK",
		Body:         body,
	}
}
