package localstorage

import (
	"context"
	"sync"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

type AreaLocalStorage struct {
	areas map[string]*entity.Area
	mutex *sync.Mutex
}

func NewAreaLocalStorage() *AreaLocalStorage {
	return &AreaLocalStorage{
		areas: make(map[string]*entity.Area),
		mutex: new(sync.Mutex),
	}
}

func (s *AreaLocalStorage) CreateArea(ctx context.Context, area *entity.Area) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.areas[area.ID] = area
	return nil
}
