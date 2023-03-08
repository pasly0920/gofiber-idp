package main

import (
	"gofiber-idp/server"
	"gofiber-idp/server/logger"
)

func main() {
	server := server.NewServer(8081, logger.Debug)

	server.Start()
}
