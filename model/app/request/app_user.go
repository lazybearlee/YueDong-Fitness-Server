package apprequest

// UserSettingReq struct
// 不支持上传头像，需要单独处理
type UserSettingReq struct {
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}
