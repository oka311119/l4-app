package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/area/repository/mock"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/uuidgen"
)

func TestAreaFlow(t *testing.T) {
	repo := new(mock.AreaStorageMock)
	uc := NewAreaUseCase(repo, &uuidgen.MockUUID{})

	var (
		userID = uc.uuidgen.V4()
		name   = "area"
		area   = entity.NewArea(
			userID,
			name,
			"default",
		)
		ctx = context.Background()
	)

	// Create Default Area
	// repo.On("CreateDefaultArea", area).Return(nil)
	// err := uc.CreateDefaultArea(ctx, userID)
	// assert.NoError(t, err)

	// Create Area
	repo.On("CreateArea", ctx, area).Return(nil)
	err := uc.CreateArea(ctx, userID, name)
	assert.NoError(t, err)
}
