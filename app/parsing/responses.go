package parsing

import "fmt"

const CLRF = "\r\n"

type Response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
}

func (receiver Response) ToBytes() []byte {
	res := fmt.Sprintf("%s %s %s%s%s",
		receiver.Version,
		receiver.StatusCode,
		receiver.ReasonPhrase,
		CLRF,
		CLRF,
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

func GetOk() Response {
	return Response{
		Version:      "HTTP/1.1",
		StatusCode:   "200",
		ReasonPhrase: "OK",
	}
}
