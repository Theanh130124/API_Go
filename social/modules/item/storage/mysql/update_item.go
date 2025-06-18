package mysql

import (
	"context"
	"social/modules/item/entity"
)

func (s *sqlStorage) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *entity.TodoItemUpdate) error {
	//no tu hieu Where("id = ?" ,id) voi cond = "id" : id
	if err := s.db.Where(cond).Updates(&dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
