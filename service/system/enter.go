package sysservice

type ServiceGroup struct {
	JwtService
	UserService
	CasbinService
	AuthorityService
	EmailService
	FileService
}
