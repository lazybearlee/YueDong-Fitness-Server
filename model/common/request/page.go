package request

import "gorm.io/gorm"

// PageInfo structure for paging
type PageInfo struct {
	Page     int    `json:"page" form:"page" query:"page"`             // Page number
	PageSize int    `json:"pageSize" form:"pageSize" query:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword" query:"keyword"`    // 用于搜索
}

// Paginate 用于调用DB进行分页，设置Offset和Limit
func (r *PageInfo) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Page <= 0 {
			r.Page = 1
		}
		switch {
		case r.PageSize > 100:
			r.PageSize = 100 // 最大100条
		case r.PageSize <= 0:
			r.PageSize = 10 // 默认10条
		}
		offset := (r.Page - 1) * r.PageSize
		return db.Offset(offset).Limit(r.PageSize)
	}
}
