package storage

import (
	"context"
	"mymodule/module/item/model"
)

func (s *sqlStore) Create(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
