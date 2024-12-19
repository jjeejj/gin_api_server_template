package api

import (
	"gin_api_server_template/app"
	userReq "gin_api_server_template/app/request/user"
	"gin_api_server_template/app/response"
	"gin_api_server_template/app/response/user"

	_struct "gin_api_server_template/app/struct"
	_const "gin_api_server_template/internal/const"
	"gin_api_server_template/internal/logger"

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

// 获取登录用户
// @Summary 获取登录用户
// @Tags user
// @Accept json
// @Produce json
// @Param GetUserInfo body user.GetUserInfoReq true "获取登录用户"
// @Param request_id header string true "请求ID"
// @Success 200 {object} user.GetUserInfoResp "成功"
// @Router /user/info [post]
func (u *userApi) GetUserInfo(c *gin.Context) {
	// 获取用户 token
	userToken, ok := c.Get(_const.CtxUserTokenKey)
	if !ok {
		response.Unauthorized(c)
		return
	}
	var userTokenS *_struct.UserToken
	if userTokenS, ok = userToken.(*_struct.UserToken); !ok {
		response.Unauthorized(c)
		return
	}
	userInfo, err := app.GA.UserInfoRepo.GetBySysId(c, userTokenS.UserSysId)
	if err != nil {
		logger.ErrorfCtx(c, "get user info err: %v", err)
		response.InternalErr(c)
		return
	}
	response.Success(c, &user.GetUserInfoResp{
		Avatar:       userInfo.Avatar,
		Birthday:     userInfo.Birthday,
		Introduction: userInfo.Introduction,
		NickName:     userInfo.NickName,
		Sex:          userInfo.Sex,
		UserSysId:    userInfo.SysId,
	})
}
