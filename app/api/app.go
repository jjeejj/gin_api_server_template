package api

import (
	"gin_api_server_template/app/response"

	"github.com/gin-gonic/gin"
)

type appApi struct{}

// @Summary app test 接口
// @Tags app
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /app/test [get]
func (app *appApi) Test(c *gin.Context) {
	response.Success(c, "test")
}
