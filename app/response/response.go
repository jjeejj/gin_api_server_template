package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success returns a success Response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error returns an error Response
func Error(c *gin.Context, code int, message string) {
	if message == "" {
		message = ErrMsgMap[code]
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// 参数错误
// 可以自定义错误信息
func ParamErr(c *gin.Context, msg string) {
	errMsg := ErrMsgMap[ErrCode_ParamErr]
	if msg != "" {
		errMsg = msg
	}
	c.JSON(http.StatusOK, Response{
		Code:    ErrCode_ParamErr,
		Message: errMsg,
		Data:    nil,
	})
}

// InternalErr 内部错误
func InternalErr(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    ErrCode_InteralErr,
		Message: ErrMsgMap[ErrCode_InteralErr],
		Data:    nil,
	})
}

// Unauthorized 认证失败
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    ErrCode_UnAuthErr,
		Message: ErrMsgMap[ErrCode_UnAuthErr],
		Data:    nil,
	})
}
