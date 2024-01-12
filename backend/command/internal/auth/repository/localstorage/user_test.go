package localstorage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/auth"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

func TestUserLocalStorage(t *testing.T) {
	s := NewUserLocalStorage()

	id1 := "id"
	user := &entity.User {
		ID: id1,
		Username: "user",
		Password: "password",
	}

	err := s.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := s.GetUser(context.Background(), "user")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	_, err = s.GetUser(context.Background(), "nonexistentuser")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
