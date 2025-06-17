package mysql

import (
	"context"
	"social/modules/item/entity"
)

func (s *sqlStorage) CreateItem(ctx context.Context, data *entity.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil { //Create truyền dữ liệu phải là &data (rút con trỏ để dữ liêu tác động thật vào db)
		return err
	}
	return nil

}
