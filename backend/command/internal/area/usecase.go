package area

import "context"

type AreaUseCase interface {
	CreateDefaultArea(ctx context.Context) error
	CreateArea(ctx context.Context, name string) error
}