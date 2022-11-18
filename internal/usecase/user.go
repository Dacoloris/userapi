package usecase

import (
	"context"
	"refactoring/domain"
	"refactoring/internal/handler/requests"
)

func (s *service) GetUsers(ctx context.Context) (domain.UserList, error) {
	return s.storage.GetUsers(ctx)
}

func (s *service) CreateUser(ctx context.Context, request requests.CreateUserRequest) (string, error) {
	return s.storage.CreateUser(ctx, request.DisplayName, request.Email)
}

func (s *service) GetUser(ctx context.Context, id string) (domain.User, error) {
	return s.storage.GetUser(ctx, id)
}

func (s *service) UpdateUserName(ctx context.Context, id, newName string) error {
	return s.storage.UpdateUserName(ctx, id, newName)
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	return s.storage.DeleteUser(ctx, id)
}
