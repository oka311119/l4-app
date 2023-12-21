package area

import (
	"context"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

type Repository interface {
	CreateArea(ctx context.Context, area *entity.Area) error
}