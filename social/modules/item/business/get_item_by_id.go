package business

import (
	"context"
	"social/modules/item/entity"
)

//-> Code thi viet bussiness truoc roi viet storage (Co chu I la de ke thua interface roi)

// Biz la bussiness
// Vi bussiness lam viec voi storage thong qua interface no
type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) //Tra ve TodoItem hoac error
}

// Lam viec voi interface
type getItemBiz struct {
	store GetItemStorage
}

// Xu ly du lieu truoc khi goi xuong storage (de ben transport goi)
func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*entity.TodoItem, error) {

	//Thuc hien create xuong entity
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err // data= null , va loi
	}
	return data, nil
}

// goi store (tu transport goi ve)
func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}
