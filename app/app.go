package app

import (
	"gin_api_server_template/app/repository/user"
	"gin_api_server_template/app/service"
	"gin_api_server_template/internal/global"
)

type GinApi struct {
	UserInfoRepo user.UserInfoRepo
	UserService  *service.UserService // 用户服务
}

var GA = &GinApi{}

// Start 初始化对应的对象
func Start() {
	// Repo 层
	GA.UserInfoRepo = user.NewDbUserInfoRepo(global.Db)

	// Service 层                                        // 应用服务
	GA.UserService = service.NewUserService(GA.UserInfoRepo) // 用户服务

	// 路由层
}
