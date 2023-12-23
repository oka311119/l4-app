package area

import "context"

type UseCase interface {
	CreateDefaultArea(ctx context.Context) error
	CreateArea(ctx context.Context, name string) error
}