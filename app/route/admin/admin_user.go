package admin

import (
	"gin_api_server_template/app/api"

	"github.com/gin-gonic/gin"
)

func genAdminUserRouter(gr *gin.RouterGroup) {
	gr.POST("list", api.AdminApi.ListUser)
}
