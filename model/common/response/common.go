package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode = 1 << 0
	ErrorCode   = 1 << 1
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse(code int, data interface{}, message string) *Response {
	return &Response{
		Code:    code,
		Data:    data,
		Message: message,
	}
}

// WriteJSON 写入JSON响应
func (r *Response) WriteJSON(c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

// Success 成功响应
func Success(c *gin.Context) {
	NewResponse(SuccessCode, nil, "操作成功").WriteJSON(c)
}

// SuccessWithData 成功响应并携带数据
func SuccessWithData(data interface{}, c *gin.Context) {
	NewResponse(SuccessCode, data, "操作成功").WriteJSON(c)
}

// SuccessWithMessage 成功响应并携带消息
func SuccessWithMessage(message string, c *gin.Context) {
	NewResponse(SuccessCode, nil, message).WriteJSON(c)
}

// SuccessWithDetailed 成功响应并携带数据和消息
func SuccessWithDetailed(data interface{}, message string, c *gin.Context) {
	NewResponse(SuccessCode, data, message).WriteJSON(c)
}

// Error 失败响应
func Error(c *gin.Context) {
	NewResponse(ErrorCode, nil, "操作失败").WriteJSON(c)
}

// ErrorWithMessage 失败响应并携带消息
func ErrorWithMessage(message string, c *gin.Context) {
	NewResponse(ErrorCode, nil, message).WriteJSON(c)
}

// ErrorWithDetailed 失败响应并携带数据和消息
func ErrorWithDetailed(data interface{}, message string, c *gin.Context) {
	NewResponse(ErrorCode, data, message).WriteJSON(c)
}

// NoAuth 无权限响应
func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, NewResponse(ErrorCode, nil, message))
}
