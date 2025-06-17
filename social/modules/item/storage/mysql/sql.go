package mysql

import (
	"gorm.io/gorm"
)

// Do minh khong viet hoa nen can NewSQLStorage de ben ngoai dung
type sqlStorage struct {
	db *gorm.DB
}

func NewSQLStorage(db *gorm.DB) *sqlStorage {
	return &sqlStorage{db: db}
}
