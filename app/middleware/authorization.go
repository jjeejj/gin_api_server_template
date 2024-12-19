package middleware

import (
	"gin_api_server_template/app/response"
	_struct "gin_api_server_template/app/struct"
	_const "gin_api_server_template/internal/const"
	"gin_api_server_template/internal/global"
	"gin_api_server_template/internal/logger"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

// Authorization login token check
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 请求头 token
		token := c.GetHeader("X-Token")
		logger.WarnfCtx(c, "req header token not exist")
		if token == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		// 从 redis 获取 token 信息
		tokenInfoStr, err := global.Rdb.Get(c, token)
		if err != nil {
			logger.ErrorfCtx(c, "get redis token info err: %v")
			response.Unauthorized(c)
			c.Abort()
			return
		}
		// 反序列化 token 信息
		userToken := &_struct.UserToken{}
		err = sonic.UnmarshalString(tokenInfoStr, userToken)
		if err != nil {
			response.ParamErr(c, "")
			c.Abort()
			return
		}
		logger.Debugf("tokenInfo: %v", userToken)
		c.Set(_const.CtxUserTokenKey, userToken)
		// 赋值调用连中的 user_id
		trace, ok := c.Value(_const.TraceCtxKey).(*logger.Trace)
		if ok {
			trace.UserSysId = userToken.UserSysId
			c.Set(_const.TraceCtxKey, trace)
		}
		c.Next()
		// token 续期
	}
}
