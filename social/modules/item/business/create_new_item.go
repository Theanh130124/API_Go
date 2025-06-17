package business

import (
	"context"
	"social/modules/item/entity"
)

// Biz la bussiness
// Vi bussiness lam viec voi storage thong qua interface no
type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *entity.TodoItemCreation) error
}
type createItemBiz struct {
	store CreateItemStorage
}
