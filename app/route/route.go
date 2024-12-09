package route

import (
	docs "gin_api_server_template/app/docs"
	"gin_api_server_template/app/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type AppRouter struct {
}

func (a *AppRouter) AddRoute(e *gin.Engine) {
	// 定义swagger文档基础信息
	docs.SwaggerInfo.Title = "gin_api_server_template"
	docs.SwaggerInfo.Description = `
	服务 api 接口 返回统一结构：
		{
			"code": 0,
			"message": "success",
			"data": {}
		}
		错误码：
			ErrCode_Success              = 0 // 成功
			ErrCode_InternalErr          = -1 // 服务内部错误
			ErrCode_ParamErr             = 100000 // 参数错误
			ErrCode_UnAuthErr            = 100001 // 未授权
	`
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 中间件
	e.Use(middleware.RequestLog())
	e.Use(middleware.Cors())
	e.Use(gin.Recovery())

	baseGroup := e.Group("/api/v1")
	genAppRouter(baseGroup.Group("/app"))

	// 用户相关的 api
	genUserRouter(baseGroup.Group("/user"))
}

func New() *AppRouter {
	return &AppRouter{}
}
