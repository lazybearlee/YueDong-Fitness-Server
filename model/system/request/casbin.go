package sysrequest

// CasbinApiInfo is a structure that defines the request body of a casbin policy request
type CasbinApiInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// CasbinInReceive is a structure that defines the request body of a casbin policy request
type CasbinInReceive struct {
	AuthorityId uint            `json:"authorityId"` // 权限id
	CasbinInfos []CasbinApiInfo `json:"casbinInfos"`
}
