package handler

import (
	"github.com/denysetiawan28/go-log/src/server/container"
)

type Handler struct {
}

func InitializeHandler(container *container.DefaultContainer) *Handler {
	return &Handler{}
}
