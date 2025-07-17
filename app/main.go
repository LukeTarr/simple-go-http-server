package main

import (
	"github.com/LukeTarr/simple-go-http-server/app/parsing/request"
	"github.com/LukeTarr/simple-go-http-server/app/parsing/response"
	"github.com/LukeTarr/simple-go-http-server/app/server"
)

func main() {
	app := server.NewServer(8080, "localhost")

	app.Get("/test", func(request request.Request) *response.Response {
		return response.GetOk("test")
	})

	app.Get("/user-agent", func(request request.Request) *response.Response {
		return response.GetOk(request.Headers["User-Agent"])
	})

	app.Get("/auth-check", func(request request.Request) *response.Response {
		return response.GetUnauthorized("Unauthorized")
	})

	app.Post("/echo", func(request request.Request) *response.Response {
		return response.GetOk(request.Body)
	})

	app.Listen()
}
