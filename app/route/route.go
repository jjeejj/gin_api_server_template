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
	docs.SwaggerInfo.Title = "sogin_api_server_template"
	docs.SwaggerInfo.Description = "服务 api 接口 \n { \"code\": 0, \"message\": \"success\", \"data\": {\"id\": 1, \"name\": \"test\"}}\"}"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "dev: http://127.0.0.1:8080"
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
