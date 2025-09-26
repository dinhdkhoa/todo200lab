package storage

import (
	"context"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Delete(cond).Error; err != nil {
		return err
	}

	return nil
}
