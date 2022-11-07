package main

import (
	"github.com/denysetiawan28/go-log/src/server"
	"github.com/denysetiawan28/go-log/src/server/container"
)

func main() {
	server.StartHttpServer(container.IntializeContainer())
}
