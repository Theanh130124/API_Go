package mysql

import (
	"context"
	"gorm.io/gorm"
	"social/common"
	"social/modules/item/entity"
)

// Storage xu ly du lieu xuong csdl (dc roi tu biz goi qua)
func (s *sqlStorage) GetItem(ctx context.Context, cond map[string]interface{}) (*entity.TodoItem, error) {
	var data entity.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil { //Create truyền dữ liệu phải là &data (rút con trỏ để dữ liêu tác động thật vào db)

		if err == gorm.ErrRecordNotFound { // lỗi không tìm thấy phải phù hợp với gorm
			return nil, common.RecordNotFound(err)
		}
		return nil, common.ErrDB(err) //Lỗi db hay gì đó
	}
	//tra ve con tro
	return &data, nil

}
