package area

import "context"

type UseCase interface {
	CreateDefaultArea(ctx context.Context, userID string) error
	CreateArea(ctx context.Context, userID string, name string) error
}