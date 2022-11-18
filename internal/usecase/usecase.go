package usecase

import (
	"context"
	"refactoring/domain"

	"go.uber.org/zap"
)

type storage interface {
	GetUsers(ctx context.Context) (domain.UserList, error)
	CreateUser(ctx context.Context, displayName, email string) (string, error)
	GetUser(ctx context.Context, id string) (domain.User, error)
	UpdateUserName(ctx context.Context, id, newName string) error
	DeleteUser(ctx context.Context, id string) error
}

type service struct {
	logger  *zap.Logger
	storage storage
}

func New(logger *zap.Logger, storage storage) *service {
	return &service{
		logger:  logger,
		storage: storage,
	}
}
