package admin

import (
	"github.com/gin-gonic/gin"
)

func GenAdminRouter(gr *gin.RouterGroup) {
	genAdminUserRouter(gr.Group("user"))
}
