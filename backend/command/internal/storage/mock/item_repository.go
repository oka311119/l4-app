package mock
import (
	"github.com/oka311119/l4-app/backend/command/internal/model"
	"github.com/oka311119/l4-app/backend/command/internal/storage"
)

type ItemRepository struct {}

func NewItemRepository() storage.ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) CreateItem(item *model.Item) error {
	return nil
}

func (r *ItemRepository) UpdateItem(item *model.Item) error {
	return nil
}
