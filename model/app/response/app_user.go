package appresponse

type UserInfo struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	HeaderImg string `json:"headerImg"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
