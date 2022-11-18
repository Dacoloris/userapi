package storage

import (
	"encoding/json"
	"io/fs"
	"os"
	"refactoring/domain"
	"time"
)

func readUserStore(fileName string) (domain.UserStore, error) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		return domain.UserStore{}, err
	}

	us := domain.UserStore{}
	err = json.Unmarshal(f, &us)
	if err != nil {
		return domain.UserStore{}, err
	}

	return us, nil
}

func saveUserStore(fileName string, us domain.UserStore) error {
	b, err := json.Marshal(us)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, b, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func newUser(displayName, email string) domain.User {
	return domain.User{
		CreatedAt:   time.Now().Round(time.Second),
		DisplayName: displayName,
		Email:       email,
	}
}
