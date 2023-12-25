package usecase

import (
	"context"
	"time"

	"github.com/oka311119/l4-app/backend/command/internal/area"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/uuidgen"
)

type AreaUseCase struct {
	areaRepo area.Repository
	uuidgen uuidgen.UUIDGenerator
}

func NewAreaUseCase(
	areaRepo area.Repository,
	uuidgen uuidgen.UUIDGenerator,
) *AreaUseCase {
	return &AreaUseCase{
		areaRepo: areaRepo,
		uuidgen: uuidgen,
	}
}

func (a *AreaUseCase) CreateDefaultArea(ctx context.Context) error {
	area := entity.NewArea(
		a.uuidgen.V4(),
		a.uuidgen.V4(),	//userid
		"default",
		time.Now(),
	) 

	return a.areaRepo.CreateArea(ctx, area)
}