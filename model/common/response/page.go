package response

// PageResponse 分页响应
type PageResponse struct {
	List     interface{} `json:"list"`     // 列表数据
	Total    int64       `json:"total"`    // 总数
	Page     int         `json:"page"`     // 当前页
	PageSize int         `json:"pageSize"` // 每页大小
}
