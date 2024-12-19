package task

import (
	"gin_api_server_template/app/service"
	"gin_api_server_template/internal/logger"

	"github.com/robfig/cron/v3"
)

// InitTask 初始化定时任务
// 如果需要其他依赖 就注入进去
func InitTask(cron *cron.Cron, userService *service.UserService) {
	cron.AddFunc("@every 1m", func() {
		logger.Debugf("cron test run")
	})
}
