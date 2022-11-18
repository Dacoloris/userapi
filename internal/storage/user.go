package storage

import (
	"context"
	"refactoring/domain"
	"strconv"
)

func (s *storage) GetUsers(_ context.Context) (domain.UserList, error) {
	us, err := readUserStore(s.fileName)
	if err != nil {
		return domain.UserList{}, err
	}

	return us.List, nil
}

func (s *storage) CreateUser(_ context.Context, displayName, email string) (string, error) {
	us, err := readUserStore(s.fileName)
	if err != nil {
		return "", err
	}

	us.Increment++

	id := strconv.Itoa(us.Increment)
	us.List[id] = newUser(displayName, email)

	err = saveUserStore(s.fileName, us)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *storage) GetUser(_ context.Context, id string) (domain.User, error) {
	us, err := readUserStore(s.fileName)
	if err != nil {
		return domain.User{}, err
	}

	user, ok := us.List[id]
	if !ok {
		return domain.User{}, domain.UserNotFound
	}

	return user, nil
}

func (s *storage) UpdateUserName(_ context.Context, id string, newName string) error {
	us, err := readUserStore(s.fileName)
	if err != nil {
		return err
	}

	user, ok := us.List[id]
	if !ok {
		return domain.UserNotFound
	}

	user.DisplayName = newName
	us.List[id] = user

	return saveUserStore(s.fileName, us)
}

func (s *storage) DeleteUser(_ context.Context, id string) error {
	us, err := readUserStore(s.fileName)
	if err != nil {
		return err
	}

	originalSize := len(us.List)
	delete(us.List, id)

	if len(us.List) == originalSize {
		return nil
	}

	return saveUserStore(s.fileName, us)
}
