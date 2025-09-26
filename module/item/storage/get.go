package storage

import (
	"context"
	"mymodule/module/item/model"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var item model.TodoItem

	if err := s.db.Where(cond).First(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
