package mysql

import (
	"context"

	"social/modules/item/entity"
)

func (s *sqlStorage) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(entity.TodoItem{}.TableName()).Where(cond).
		Updates(map[string]interface{}{"status": entity.ItemStatusDeleted}).
		Error; err != nil { //db.First (tim kiem 1 hang theo id)
		return err
	}
	return nil
}
