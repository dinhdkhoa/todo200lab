package uc

import (
	"context"
	"mymodule/module/item/model"

	"github.com/google/uuid"
)

type updateItemUseCase struct {
	store UpdateItemStorage
}

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
}

func NewUpdateItemUseCase(s UpdateItemStorage) *updateItemUseCase {
	return &updateItemUseCase{store: s}
}

func (uc *updateItemUseCase) UpdateItemByIdUC(ctx context.Context, data *model.TodoItemUpdate, id uuid.UUID) error {
	item, err := uc.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if item.Status == "Deleted" {
		return model.ErrorItemIsDeleted
	}

	if err := uc.store.UpdateItem(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
