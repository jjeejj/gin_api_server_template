package middleware

import (
	"bytes"
	"gin_api_server_template/internal/global"
	"io"

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
		c.Writer = crw
		reqBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		global.Logger.Infof("Request: %s %s %s", c.Request.Method, c.Request.RequestURI, reqBody)
		c.Next()
		// 在请求处理之后记录日志
		respBody := crw.body.String()
		global.Logger.Infof("Response: %s %s %s", c.Request.Method, c.Request.RequestURI, respBody)
	}
}
