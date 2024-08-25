package sysrequest

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v4"
)

// BaseClaims Base claims structure
type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}
