package bootstrap

import (
	"fmt"
	"gin_api_server_template/internal/global"
	"gin_api_server_template/internal/logger"
	"gin_api_server_template/internal/mysql"
	"gin_api_server_template/internal/redis"
	"os"

	"gin_api_server_template/internal/config"
)

// 初始化服务启动相关的配置
func init() {
	var err error
	global.Config = config.GetConfig()
	global.RootDir, _ = os.Getwd()
	// 初始化日志
	err = logger.Init(global.Config.Log, global.RootDir)
	if err != nil {
		panic(err)
	}
	// mysql 配置必须传
	if global.Config.Mysql.Host == "" || global.Config.Mysql.Port == "" {
		panic("mysql config is empty")
	}
	global.Db = mysql.New(
		&mysql.MysqlConfig{
			Host:     global.Config.Mysql.Host,
			Port:     global.Config.Mysql.Port,
			User:     global.Config.Mysql.User,
			Password: global.Config.Mysql.Password,
			Database: global.Config.Mysql.Database,
			Charset:  global.Config.Mysql.Host,
		},
	)
	global.Rdb = redis.New(&redis.RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Password,
		Db:       global.Config.Redis.Db,
		Prefix:   global.Config.Redis.Prefix,
	})
}
