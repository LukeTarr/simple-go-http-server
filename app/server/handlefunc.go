package server

import (
	"github.com/LukeTarr/simple-go-http-server/app/parsing/request"
	"github.com/LukeTarr/simple-go-http-server/app/parsing/response"
)

type FunctionPath struct {
	Path            string
	Method          string
	HandlerFunction func(request request.Request) *response.Response
}
