package storage

import (
	"github.com/oka311119/l4-app/backend/command/internal/model"
)

type ItemRepository interface {
	CreateItem(item *model.Item) error
	UpdateItem(item *model.Item) error
}