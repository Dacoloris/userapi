package handler

import (
	"context"
	"refactoring/domain"
	"refactoring/internal/handler/requests"

	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	useCase useCase
}

type useCase interface {
	GetUsers(ctx context.Context) (domain.UserList, error)
	CreateUser(ctx context.Context, request requests.CreateUserRequest) (string, error)
	GetUser(ctx context.Context, id string) (domain.User, error)
	UpdateUserName(ctx context.Context, id, newName string) error
	DeleteUser(ctx context.Context, id string) error
}

func New(logger *zap.Logger, useCase useCase) *Handler {
	return &Handler{
		logger:  logger,
		useCase: useCase,
	}
}
