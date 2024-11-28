package _struct

const (
	DefaultPageSize = 10   // 默认每页大小
	DefaultPage     = 1    // 默认第几页
	MaxPageSize     = 1000 // 每页的最大数量
)

// Page 信息
type Page struct {
	// 第几页，从 1 开始
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// Offset
	Offset int `json:"offset"`
	// Limit
	Limit int `json:"limit"`
}

func NewPage(page int, pageSize int) *Page {
	return &Page{
		Page:     page,
		PageSize: pageSize,
		Offset:   (page - 1) * pageSize,
		Limit:    pageSize,
	}
}
