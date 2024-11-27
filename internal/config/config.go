package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

// Config 配置信息
type Config struct {
	App   AppConfig
	Log   LogConfig
	Mysql MysqlConfig
	Redis RedisConfig
}

// AppConfig 应用配置
type AppConfig struct {
	Env        string
	Port       string
	Name       string
	IsTaskNode bool `json:"is_task_node" mapstructure:"is_task_node"`
}

// MysqlConfig mysql配置
type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Charset  string
}

// redis配置
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
	Prefix   string
}

// log 配置
type LogConfig struct {
	FilePath         string `json:"file_path" mapstructure:"file_path"`
	FileName         string `json:"file_name" mapstructure:"file_name"`
	Level            string `json:"level" mapstructure:"level"`
	ConsoleEnable    bool   `json:"console_enable" mapstructure:"console_enable"`
	ConsoleOrmEnable bool   `json:"console_orm_enable" mapstructure:"console_orm_enable"`
	IsCompress       bool   `json:"is_compress" mapstructure:"is_compress"` // 是否压缩
	MaxAge           int    `json:"max_age" mapstructure:"max_age"`         // 日志文件保留天数
	MaxBackups       int    `json:"max_backups" mapstructure:"max_backups"` // 最大保留文件数量
	MaxSize          int    `json:"max_size" mapstructure:"max_size"`       // 日志文件最大大小
}

// GetConfig 从配置文件获取配置信息
func GetConfig() (c *Config) {
	// 获取项目运营的跟目录
	// _, filePath, _, ok := runtime.Caller(0) // 获取当前文件路径
	// if !ok {
	// 	panic("can not get current file info")
	// }
	// dirPath := path.Dir(filePath)
	// fmt.Printf("dirPath: %s\n", dirPath)
	// rootDir := path.Join(dirPath, "../..")
	rootDir, _ := os.Getwd()
	// fmt.Printf("rootDir: %s\n", rootDir)

	// 读取配置文件
	v := viper.New()
	v.SetConfigFile(path.Join(rootDir, "config.yaml"))
	// v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	return
}
