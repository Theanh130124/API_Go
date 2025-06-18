package mysql

import (
	"context"
	"social/common"
	"social/modules/item/entity"
)

// Storage xu ly du lieu xuong csdl (dc roi tu biz goi qua)
func (s *sqlStorage) ListItem(ctx context.Context,
	filter *entity.Filter, // truyen con tro thi co the nil cho truong du lieu nay (con neu khong truyen thi filter luon phai truyen vao)
	paging *common.Paging,
	moreKey ...string) ([]entity.TodoItem, error) { //Tra ve TodoItem hoac error {
	//Nếu không truyền gì mặc định là page 1 và limit = 10
	var res []entity.TodoItem // là các slide Todo( hay là List Item)

	//Do cái deleteShort vẫn còn nên mình cần lọc ra

	// <> là khác
	db := s.db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	//Find để tìm nhiều dòng dữ liệu và Order theo id desc (giảm -> mới trc)

	if err := db.Table(entity.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&res).Error; err != nil { //Create truyền dữ liệu phải là &data (rút con trỏ để dữ liêu tác động thật vào db)

		return nil, err
	}
	return res, nil
}
