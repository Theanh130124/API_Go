package common

// 1 lưu ý là tên các attribute của struct đều phải viết hoa chữ đầu , để truyền json đc

// interface là truyền cái gì cũng được
type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes { // muốn thực sự thây đổi phải truyền con trỏ
	return &successRes{Data: data, Paging: paging, Filter: filter} // muốn thực sự  phải thực hiện trên giá trị con trỏ &successRes
	//successRes chỉ là địa chỉ con trỏ
}

// Đôi khi chỉ cần truyền data
func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data, Paging: nil, Filter: nil}
}
