package common

import "time"

type SQLModel struct {
	Id        int       `json:"id" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`            //hiện time null json nếu không có
	UpdatedAt time.Time `json:"updated_at ,omitempty" gorm:"column:updated_at"` // omitempty bỏ nil  , bỏ false , bỏ qua chuoi rong
}
