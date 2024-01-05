package usecase

import (
    "context"
    "testing"

	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/uuidgen"
)

func TestAreaFlow(t *testing.T) {
    repo := new(mock.AreaStorageMock)
    uc := NewAreaUseCase(repo, &uuidgen.MockUUID{})

    var (
        areaID = uc.uuidgen.V4()
        userID = uc.uuidgen.V4()
        name = "area-name"
        area = entity.NewArea(
            areaID,
            userID,
            "default",
        )
        ctx = context.Background()
    )

    // Create Default Area
    repo.On("CreateDefaultArea", area).Return(nil)
    err := uc.CreateDefaultArea(ctx, userID)
    assert.NoError(t, err)

    // Create Area
    repo.On("CreateArea", area).Return(nil)
    err := uc.CreateArea(ctx, userID, name)
    assert.NoError(t, err)
}
