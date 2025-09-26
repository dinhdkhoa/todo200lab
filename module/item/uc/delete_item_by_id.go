package uc

import (
	"context"
	"mymodule/module/item/model"

	"github.com/google/uuid"
)

type deleteItemUseCase struct {
	store deleteItemStorage
}

type deleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

func NewDeleteItemUseCase(s deleteItemStorage) *deleteItemUseCase {
	return &deleteItemUseCase{store: s}
}

func (uc *deleteItemUseCase) DeleteItemByIdUC(ctx context.Context, id uuid.UUID) error {
	item, err := uc.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if item.Status == "Deleted" {
		return model.ErrorItemIsDeleted
	}

	if err := uc.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
