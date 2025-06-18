package business

import (
	"context"
	"social/modules/item/entity"
)

//-> Code thi viet bussiness truoc roi viet storage (Co chu I la de ke thua interface roi)

// Biz la bussiness
// Vi bussiness lam viec voi storage thong qua interface no  (Shift +F6 de thay doi nhieu cho 1 luc)
type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) //Tra ve TodoItem hoac error
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *entity.TodoItemUpdate) error
}

// Lam viec voi interface
type updateItemBiz struct {
	store UpdateItemStorage
}

// Xu ly du lieu truoc khi goi xuong storage (de ben transport goi)
func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *entity.TodoItemUpdate) error {

	//Minh can get de lay id de xoa
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err // data= null , va loi
	}

	//khong update data khong co status va deleted
	if data.Status != nil && *data.Status == entity.ItemStatusDeleted {
		return entity.ErrItemDeleted
	}
	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}

// goi store (tu transport goi ve)
func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: store}
}
