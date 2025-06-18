package business

import (
	"context"
	"social/common"
	"social/modules/item/entity"
)

//-> Code thi viet bussiness truoc roi viet storage (Co chu I la de ke thua interface roi)

// Biz la bussiness
// Vi bussiness lam viec voi storage thong qua interface no        ...string  la mang string co hay khong cung duoc
type ListItemStorage interface {
	ListItem(ctx context.Context,
		filter *entity.Filter, // truyen con tro thi co the nil cho truong du lieu nay (con neu khong truyen thi filter luon phai truyen vao)
		paging *common.Paging,
		moreKey ...string) ([]entity.TodoItem, error) //Tra ve TodoItem hoac error
}

// Lam viec voi interface
type listItemBiz struct {
	store ListItemStorage
}

// Xu ly du lieu truoc khi goi xuong storage (de ben transport goi)
func (biz *listItemBiz) ListItem(ctx context.Context,
	filter *entity.Filter,
	paging *common.Paging) ([]entity.TodoItem, error) {

	//Thuc hien create xuong entity
	data, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, err // data= null , va loi
	}
	return data, nil
}

// goi store (tu transport goi ve)
func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}
