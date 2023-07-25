package handler

import (
	"golangtest/src/server/container"
	"golangtest/src/usecase"
)

type Handler struct {
	useCase usecase.UseCase
}

func InitializeHandler(container *container.DefaultContainer) *Handler {
	return &Handler{
		useCase: container.Usecase,
	}
}
