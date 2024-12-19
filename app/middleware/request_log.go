package middleware

import (
	"bytes"
	"gin_api_server_template/internal/logger"
	"io"

	_const "gin_api_server_template/internal/const"

	"github.com/gin-gonic/gin"
)

// CustomResponseWriter 封装 gin ResponseWriter 用于获取回包内容。
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLog 打印请求日志
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求处理之前记录日志
		// / 使用自定义 ResponseWriter
		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		// 获取请求头的 request_id
		requestId := c.GetHeader("request_id")
		c.Set(_const.TraceCtxKey, &logger.Trace{
			TraceId: requestId,
			Host:    c.Request.Host,
			Path:    c.Request.URL.Path,
		})
		c.Writer = crw
		reqBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		logger.InfofCtx(c, "start request, method: %s, body: %sv", c.Request.Method, string(reqBody))
		c.Next()
		// 在请求处理之后记录日志
		respBody := crw.body.String()
		logger.InfofCtx(c, "end request, response_body: %s", respBody)
	}
}
