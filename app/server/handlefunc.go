package server

import "github.com/LukeTarr/simple-go-http-server/app/parsing"

type FunctionPath struct {
	Path            string
	Method          string
	HandlerFunction func(request parsing.Request) string
}
