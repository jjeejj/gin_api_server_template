package app

import (
	"gin_api_server_template/app/repository/user"
	"gin_api_server_template/app/service"
	"gin_api_server_template/app/task"
	"gin_api_server_template/internal/global"

	_const "gin_api_server_template/internal/const"

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
	// 非生产环境，自动合并 model
	if global.Config.App.Env != _const.Env_Prod {
		initModels(global.Db)
	}
	// Repo 层
	GA.UserInfoRepo = user.NewDbUserInfoRepo(global.Db)

	// Service 层                                        // 应用服务
	GA.UserService = service.NewUserService(GA.UserInfoRepo) // 用户服务

	// 路由层

	// 定时器
	GA.Cron = cron.New(cron.WithSeconds())

	// 只有执行任务的节点才执行
	if global.Config.App.IsTaskNode {
		// 添加定时任务
		task.InitTask(GA.Cron, GA.UserService)
		// 启动定时任务
		GA.Cron.Start()
	}
}
