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
		ctx = context.Background()

        areaID = uc.uuidgen.V4()
		userID = uc.uuidgen.V4()
		name   = "new-area"
        area = entity.NewArea(
            areaID,
		    userID,
		    name,
	    )
	)

	// Create Area
	repo.On("CreateArea", area).Return(nil)
    err := uc.CreateArea(ctx, userID, name)
	assert.NoError(t, err)
}
