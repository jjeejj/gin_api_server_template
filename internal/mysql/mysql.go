package mysql

import (
	"gin_api_server_template/internal/logger"

	mysqlConfig "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Charset  string
}

func New(config *MysqlConfig) (db *gorm.DB) {
	// config := global.Config.Mysql
	// dsn
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	mysqlConfig := mysql.Config{
		DSN: dsn,
		DSNConfig: &mysqlConfig.Config{
			ConnectionAttributes: "parseTime:true,loc=UTC",
		},
		DefaultStringSize:        256,  // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持

	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		// Logger: global.Config.Log,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Debugf("mysql connect success %s:%s", config.Host, config.Port)
	return

}
