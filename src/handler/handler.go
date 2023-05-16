package handler

import (
	"github.com/denysetiawan28/go-log/src/server/container"
)

type Handler struct {
	HelloHandler *helloHandler
}

func InitializeHandler(container *container.DefaultContainer) *Handler {
	return &Handler{HelloHandler: NewHelloHandler()}
}
