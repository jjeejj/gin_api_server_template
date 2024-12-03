package task

import (
	"gin_api_server_template/internal/global"

	"github.com/robfig/cron/v3"
)

func AddTask(cron *cron.Cron) {
	cron.AddFunc("@every 1m", func() {
		global.Logger.Debugf("cron test run")
	})
}
