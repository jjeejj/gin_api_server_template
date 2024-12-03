package app

import (
	"gin_api_server_template/app/repository/user"
	"gin_api_server_template/app/service"
	"gin_api_server_template/app/task"
	"gin_api_server_template/internal/global"

	"github.com/robfig/cron/v3"
)

type GinApi struct {
	UserInfoRepo user.UserInfoRepo
	UserService  *service.UserService // 用户服务

	Cron *cron.Cron
}

var GA = &GinApi{}

// Start 初始化对应的对象
func Start() {
	// Repo 层
	GA.UserInfoRepo = user.NewDbUserInfoRepo(global.Db)

	// Service 层                                        // 应用服务
	GA.UserService = service.NewUserService(GA.UserInfoRepo) // 用户服务

	// 路由层

	// 定时器
	GA.Cron = cron.New(cron.WithSeconds())

	// 添加定时任务
	task.InitTask(GA.Cron, GA.UserService)
	// 启动定时任务
	GA.Cron.Start()
}
