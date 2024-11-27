package server

import (
	"github.com/gin-gonic/gin"
)

type RouterGeneratorInterface interface {
	AddRoute(server *gin.Engine)
}
