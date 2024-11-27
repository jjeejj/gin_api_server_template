package route

import (
	"gin_api_server_template/app/api"
	"gin_api_server_template/app/middleware"

	"github.com/gin-gonic/gin"
)

func genUserRouter(gr *gin.RouterGroup) {
	// 登录接口，排除权限验证
	gr.POST("/login", api.UserApi.Login)
	// 其他接口需要权限验证
	gr.Use(middleware.Authorization())
}
