package localstorage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

func TestCreateArea(t *testing.T) {
	s := NewAreaLocalStorage()

	id1 := "id"
	area := &entity.Area{
		ID: id1,
		UserID: "id",
		Name: "area-name",
    }

    err := s.CreateArea(context.Background(), area)
	assert.NoError(t, err)
}
