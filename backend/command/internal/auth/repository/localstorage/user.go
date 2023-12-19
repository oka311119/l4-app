package localstorage

import (
	"context"
	"sync"

	"github.com/oka311119/l4-app/backend/command/internal/auth"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)


type UserLocalStorage struct {
	users map[string]*entity.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[string]*entity.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserLocalStorage) CreateUser(ctx context.Context, user *entity.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.users[user.ID] = user
	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}