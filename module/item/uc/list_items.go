package uc

import (
	"context"
	"mymodule/common"
	"mymodule/module/item/model"
)

type listItemsUseCase struct {
	store ListItemsStorage
}

type ListItemsStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filters,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

func NewListItemsUseCase(s ListItemsStorage) *listItemsUseCase {
	return &listItemsUseCase{store: s}
}

func (uc *listItemsUseCase) ListItemsUC(ctx context.Context,
	filter *model.Filters,
	paging *common.Paging,
	moreKeys ...string) ([]model.TodoItem, error) {

	items, err := uc.store.ListItem(ctx, filter, paging, moreKeys...)

	if err != nil {
		return nil, err
	}
	return items, nil
}
