package memory

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/silverspase/k8s-prod-service/internal/todo"
	"github.com/silverspase/k8s-prod-service/internal/todo/model"
)

type memoryStorage struct {
	items  map[string]model.Item // TODO change to sync.Map
	logger *zap.Logger
}

func NewMemoryStorage(logger *zap.Logger) todo.Repository {
	return &memoryStorage{
		items:  make(map[string]model.Item),
		logger: logger,
	}
}

func (m memoryStorage) CreateItem(ctx context.Context, item model.Item) (string, error) {
	item.ID = uuid.New().String()
	m.items[item.ID] = item

	return item.ID, nil
}

func (m memoryStorage) GetAllItems(ctx context.Context) (res []model.Item, err error) {
	for _, item := range m.items {
		res = append(res, item)
	}

	return res, nil
}

func (m memoryStorage) GetItem(ctx context.Context, id string) (model.Item, bool) {
	item, ok := m.items[id]
	return item, ok
}

func (m memoryStorage) UpdateItem(ctx context.Context, item model.Item) (string, error) {
	_, ok := m.items[item.ID]
	if !ok {
		return "", errors.New("item with given id not found, nothing to update")
	}

	m.items[item.ID] = item

	return item.ID, nil
}

func (m memoryStorage) DeleteItem(ctx context.Context, id string) (string, error) {
	_, ok := m.items[id]
	if !ok {
		return "", errors.New("item with given id not found, nothing to delete")
	}

	delete(m.items, id)

	return id, nil
}