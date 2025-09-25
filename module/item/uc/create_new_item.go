package uc

import (
	"context"
	"mymodule/module/item/model"
)


type createItemUseCase struct {
	store CreateItemStorage
}

type CreateItemStorage interface {
	Create(ctx context.Context, data *model.TodoItemCreation) error
}

func NewCreateItemUseCase(s CreateItemStorage) *createItemUseCase {
	return &createItemUseCase{store: s}
}

func (uc *createItemUseCase) CreateItemUC(ctx context.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := uc.store.Create(ctx, data); err != nil {
		return err
	}
	
	return nil
}