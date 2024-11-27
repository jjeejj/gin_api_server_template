package api

import (
	userReq "gin_api_server_template/app/request/user"
	"gin_api_server_template/app/response"
	"gin_api_server_template/app/response/user"

	"github.com/gin-gonic/gin"
)

type userApi struct {
}

// 登录
// @Summary 手机\邮箱验证码登录
// @Tags user
// @Accept json
// @Produce json
// @Param LoginReq body user.UserLoginReq true "用户登录"
// @Param request_id header string true "请求ID"
// @Success 200 {object} user.UserLoginResp "登录成功"
// @Router /user/login [post]
func (u *userApi) Login(c *gin.Context) {
	reqBody := &userReq.UserLoginReq{}
	if err := c.ShouldBindJSON(reqBody); err != nil {
		response.ParamErr(c, "")
		return
	}
	response.Success(c, &user.UserLoginResp{})
}
