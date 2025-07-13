package parsing

import "fmt"

const CLRF = "\\r\\n"

type Response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
}

func (receiver Response) ToBytes() []byte {
	res := fmt.Sprintf("%s %s %s %s",
		receiver.Version,
		receiver.StatusCode,
		receiver.ReasonPhrase,
		CLRF,
	)
	return []byte(res)
}
