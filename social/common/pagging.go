package common

// Phân trang  -> (query string -> ?page=1&limit=3) muốn dùng query string cần form
type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

// Xử lý default (nếu không gửi page mặc định là)
func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
}
