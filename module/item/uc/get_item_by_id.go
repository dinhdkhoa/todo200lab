package uc

import (
	"context"
	"mymodule/module/item/model"

	"github.com/google/uuid"
)

type getItemUseCase struct {
	store UpdateItemStorage
}

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

func NewGetItemUseCase(s UpdateItemStorage) *updateItemUseCase {
	return &updateItemUseCase{store: s}
}

func (uc *updateItemUseCase) GetItemByIdUC(ctx context.Context, id uuid.UUID) (*model.TodoItem, error) {
	var item *model.TodoItem
	item, err := uc.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return item, nil
}
