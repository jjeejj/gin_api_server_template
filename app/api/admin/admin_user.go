package admin

import (
	"gin_api_server_template/app/response"

	"github.com/gin-gonic/gin"
)

// 用户列表
// @Summary 用户列表
// @Tags admin
// @Accept json
// @Produce json
// @Param ListUserReq body admin.ListUserReq true "用户登录"
// @Param request_id header string true "请求ID"
// @Success 200 {object} admin.ListUserResp "登录成功"
// @Router /admin/user/list [post]
func (a *adminApi) ListUser(c *gin.Context) {
	response.Success(c, "")
}
