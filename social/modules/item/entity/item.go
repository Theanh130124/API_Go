package entity

import (
	"social/common"
)

// Phần json là tên key gửi trên postman
type TodoItem struct {
	//that ra khong can dien gorm -> vì nó đã tự map đúng name json
	common.SQLModel             //id createAt va UpdateAt
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	Status          *ItemStatus `json:"status"` //lấy * để có thể null

}

// Nhận json body từ sv vào structure (binh thuong gorm khong can ghi no tu scan)

type TodoItemCreation struct {
	//Du lieu trong postman se la cac truong nay
	Id          int         `json:"-" gorm:"id"` //Tao khong can truyen id nen ->  "-" la khong truyen gi
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       string  `json:"title"`
	Description *string `json:"description"` //vì gorm sẽ dưa vào chuỗi rỗng hoặc khác new để (bỏ qua update ) khi là con trỏ thì sẽ trỏ tới 1 giá trị new -> gorm sẽ update dù chơi chuỗi rỗng
	Status      string  `json:"status"`
}

//Dùng chung

func (TodoItem) TableName() string {
	return "todo_items" // ten bang sẽ tác động đến
}

//Hàm để GORM insert struct xuống db

// func(reciver) function_name kieuDL
func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName() // ten bang luc nay cung la todo_items
	//ToDoItem tác động lên bảng todo_items nên mình chỉ cần return TableName(vì thằng này cũng cần bảng todo_items)
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
