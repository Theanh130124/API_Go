package business

import (
	"context"
	"social/modules/item/entity"
	"strings"
)

// Biz la bussiness
// Vi bussiness lam viec voi storage thong qua interface no
type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *entity.TodoItemCreation) error
}
type createItemBiz struct {
	store CreateItemStorage
}

func (biz *createItemBiz) CreateItem(ctx context.Context, data *entity.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title) //TrimSpace bo khoang trang 2 dau
	if title == "" {                       //neu bi bo trong
		return entity.ErrTitleIsBlank //tra ve loi minh tu dinh nghia
	}

	//Thuc hien create xuong entity
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}
