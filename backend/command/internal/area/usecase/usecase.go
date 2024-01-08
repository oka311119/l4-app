package usecase

import (
	"context"

	"github.com/oka311119/l4-app/backend/command/internal/area"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/uuidgen"
)

type AreaUseCase struct {
	areaRepo area.Repository
	uuidgen  uuidgen.UUIDGenerator
}

func NewAreaUseCase(
	areaRepo area.Repository,
	uuidgen uuidgen.UUIDGenerator,
) *AreaUseCase {
	return &AreaUseCase{
		areaRepo: areaRepo,
		uuidgen:  uuidgen,
	}
}

func (a *AreaUseCase) CreateArea(ctx context.Context, userID, name string) error {
    if name == entity.DefaultAreaName {
        return area.ErrAreaIsAlreadyExists
    }

    area := entity.NewArea(
		a.uuidgen.V4(),
		userID,
		name,
	)

	return a.areaRepo.CreateArea(ctx, area)
}
