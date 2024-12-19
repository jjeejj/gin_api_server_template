package global

import (
	"gin_api_server_template/internal/config"
	"gin_api_server_template/internal/redis"

	"gorm.io/gorm"
)

var (
	Config *config.Config
	// Logger  *zap.SugaredLogger
	RootDir string   // 项目的跟目录
	Db      *gorm.DB // mysql 连接
	Rdb     *redis.RedisClient
)
