package route

import (
	"gin_api_server_template/app/api"

	"github.com/gin-gonic/gin"
)

func genAppRouter(gr *gin.RouterGroup) {
	gr.GET("/test", api.AppApi.Test)
}
