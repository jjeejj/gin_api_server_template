package logger

import (
	"context"

	_const "gin_api_server_template/internal/const"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Trace struct {
	Host      string `json:"host"` // 请求路径
	Path      string `json:"path"` // 请求路径
	TraceId   string `json:"trace_id"`
	UserSysId string `json:"user_sys_id"`
	// SpanId  string `json:"span_id"`
	// Caller     string `json:"caller"`
	// SrcMethods string `json:"src_methods,omitempty"`
}

func GetTraceCtx(c *gin.Context) context.Context {
	return c.MustGet(_const.TraceCtxKey).(context.Context)
}

func (t *Trace) ToZapFields() []zap.Field {
	return []zap.Field{
		zap.String("trace_id", t.TraceId),
		zap.String("user_sys_id", t.UserSysId),
		zap.String("host", t.Host),
		zap.String("path", t.Path),
	}
}
