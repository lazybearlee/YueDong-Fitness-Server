package response

import "github.com/lazybearlee/yuedong-fitness/model/system"

// UserResponse struct
type UserResponse struct {
	User system.SysUser `json:"user"`
}

// LoginResponse struct
type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
