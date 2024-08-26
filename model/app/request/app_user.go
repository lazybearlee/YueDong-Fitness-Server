package apprequest

// UserUpdateInfoReq struct
// 不支持上传头像，需要单独处理
type UserUpdateInfoReq struct {
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

// UserUpdatePasswordReq struct
type UserUpdatePasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
