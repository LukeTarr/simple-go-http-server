package main

import (
	"github.com/LukeTarr/simple-go-http-server/app/parsing"
	"github.com/LukeTarr/simple-go-http-server/app/server"
)

func main() {
	app := server.Server{
		Port:    8080,
		Address: "0.0.0.0",
	}

	app.Get("/test", func(request parsing.Request) string {
		return "test"
	})

	app.Get("/user-agent", func(request parsing.Request) string {
		return request.Headers["User-Agent"]
	})

	app.Post("/echo", func(request parsing.Request) string {
		return request.Body
	})

	app.Listen()
}
