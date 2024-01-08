package area

import "context"

type UseCase interface {
	CreateArea(ctx context.Context, userID string, name string) error
}
