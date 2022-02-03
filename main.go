package main

import (
	"github.com/labstack/echo/v4"
	"github.com/xImouto/sna-go/handlers"
)

var server = echo.New()

func main() {
	// supporting middlewares
	server.GET("/", handlers.Root)
	server.GET("/user", handlers.UserFind)
	server.GET("/user/:username", handlers.UserFindOne)

	// start server
	server.Logger.Fatal(server.Start(":3000"))
}
