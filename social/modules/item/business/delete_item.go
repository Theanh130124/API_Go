package business

import (
	"context"
	"social/modules/item/entity"
)

//-> Code thi viet bussiness truoc roi viet storage (Co chu I la de ke thua interface roi)

// Biz la bussiness
// Vi bussiness lam viec voi storage thong qua interface no  (Shift +F6 de thay doi nhieu cho 1 luc)
type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) //Tra ve TodoItem hoac error
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

// Lam viec voi interface
type deleteItemBiz struct {
	store DeleteItemStorage
}

// Xu ly du lieu truoc khi goi xuong storage (de ben transport goi)
func (biz *deleteItemBiz) DeleteItemById(ctx context.Context, id int) error {

	//Minh can get de lay id de xoa
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err // data= null , va loi
	}

	if data.Status != nil && *data.Status == entity.ItemStatusDeleted {
		return entity.ErrItemDeleted
	}
	if err := biz.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}

// goi store (tu transport goi ve)
func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{store: store}
}
