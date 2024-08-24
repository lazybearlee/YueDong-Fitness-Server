package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	sysrequest "github.com/lazybearlee/yuedong-fitness/model/system/request"
	"net"
	"time"
)

type JWT struct {
	SigningKey []byte // 签名密钥
}

var (
	ErrTokenExpired     = errors.New("token is expired")           // Token 过期
	ErrTokenNotValidYet = errors.New("token not active yet")       // Token 尚未激活
	ErrTokenMalformed   = errors.New("that's not even a token")    // Token 格式错误
	ErrTokenInvalid     = errors.New("couldn't handle this token") // 无法处理此 Token
)

// NewJWT 新建一个 JWT 实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.FITNESS_CONFIG.JWT.SigningKey),
	}
}

// CreateClaims 创建 JWT 的 Claims
func (j *JWT) CreateClaims(baseClaims sysrequest.BaseClaims) sysrequest.CustomClaims {
	// 首先获取缓冲时间和过期时间
	bufferTime, _ := ParseDuration(global.FITNESS_CONFIG.JWT.BufferTime)
	expiresTime, _ := ParseDuration(global.FITNESS_CONFIG.JWT.ExpiresTime)
	// 创建 CustomClaims
	claims := sysrequest.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bufferTime / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"FITNESS"},
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresTime)),
			Issuer:    global.FITNESS_CONFIG.JWT.Issuer,
		},
	}
	return claims
}

// CreateToken 创建 Token
func (j *JWT) CreateToken(claims sysrequest.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 通过旧 Token 创建新 Token，对于相同的 Token，只会生成一个新 Token
func (j *JWT) CreateTokenByOldToken(oldToken string, claims sysrequest.CustomClaims) (string, error) {
	v, err, _ := global.FITNESS_CC.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 Token，需要注意的是，如果Token非法，会返回错误
func (j *JWT) ParseToken(tokenString string) (*sysrequest.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &sysrequest.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*sysrequest.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, ErrTokenInvalid
}

// NeedRefreshToken 判断是否需要刷新 Token
func (j *JWT) NeedRefreshToken(claims *sysrequest.CustomClaims) bool {
	return claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime
}

// GetToken 获取请求中的Token
func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

// SetToken 设置Token
func SetToken(c *gin.Context, token string, maxAge int) {
	// 这里主要是获取请求的host，然后设置cookie的domain
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	// 如果是ip地址，就设置domain为空
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

// ClearToken 清除Token
func ClearToken(c *gin.Context) {
	// 这里主要是获取请求的host，然后设置cookie的domain
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	// 如果是ip地址，就设置domain为空
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

// GetJWTClaims 获取JWT中的Claims
func GetJWTClaims(c *gin.Context) (*sysrequest.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// GetCustomClaims 获取CustomClaims
func GetCustomClaims(c *gin.Context) *sysrequest.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetJWTClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*sysrequest.CustomClaims)
		return waitUse
	}
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) uint {
	customClaims := GetCustomClaims(c)
	if customClaims == nil {
		return 0
	}
	return customClaims.BaseClaims.ID
}

// GetUserUuid 获取用户Uuid
func GetUserUuid(c *gin.Context) uuid.UUID {
	customClaims := GetCustomClaims(c)
	if customClaims == nil {
		return uuid.Nil
	}
	return customClaims.UUID
}

// GetUserName 获取用户名
func GetUserName(c *gin.Context) string {
	customClaims := GetCustomClaims(c)
	if customClaims == nil {
		return ""
	}
	return customClaims.Username
}

// GetUserAuthorityId 获取用户角色ID
func GetUserAuthorityId(c *gin.Context) uint {
	customClaims := GetCustomClaims(c)
	if customClaims == nil {
		return 0
	}
	return customClaims.AuthorityId
}

// NewLoginToken 创建登录Token
func NewLoginToken(login sysmodel.Login) (string, sysrequest.CustomClaims, error) {
	j := NewJWT()
	claims := j.CreateClaims(sysrequest.BaseClaims{
		ID:          login.GetUserId(),
		UUID:        login.GetUUID(),
		NickName:    login.GetNickname(),
		Username:    login.GetUsername(),
		AuthorityId: login.GetAuthorityId(),
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", sysrequest.CustomClaims{}, err
	}
	return token, claims, nil
}
