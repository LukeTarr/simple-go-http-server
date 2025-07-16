package main

import (
	"github.com/LukeTarr/simple-go-http-server/app/server"
)

func main() {
	app := server.Server{
		Port:    8080,
		Address: "0.0.0.0",
	}

	app.Listen()
}
