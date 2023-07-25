package main

import (
	"golangtest/src/server"
	"golangtest/src/server/container"
)

func main() {
	server.StartHttpServer(container.IntializeContainer())
}
