package main

import (
	"github.com/labstack/echo/v4"
)

var server = echo.New()

func main() {
	// supporting middlewares
	server.GET('/', Root)

	// start server
	server.Logger.Fatal(server.Start(":3000"))
}
