package mysql

import (
	"context"
	"social/modules/item/entity"
)

// Storage xu ly du lieu xuong csdl (dc roi tu biz goi qua)
func (s *sqlStorage) GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	var data entity.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil { //Create truyền dữ liệu phải là &data (rút con trỏ để dữ liêu tác động thật vào db)
		return nil, err
	}
	//tra ve con tro
	return &data, nil

}
