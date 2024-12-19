package logger

import (
	"gin_api_server_template/internal/config"
	"gin_api_server_template/internal/global"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	options  []zap.Option
	confInfo config.LogConfig
	rootDir  string
)

// New 初始化日志
// configInfo 日志配置信息
// dir: 日志根目录
func new(configInfo config.LogConfig, dir string) (Logger *zap.Logger, err error) {
	confInfo = configInfo
	rootDir = dir
	logDir := path.Join(rootDir, confInfo.FilePath)
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}
	loggerConf := getConfig()
	writer, err := getWriteSyncer()
	if err != nil {
		return
	}
	coreTeeArr := []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewJSONEncoder(loggerConf.EncoderConfig),
			writer,
			loggerConf.Level,
		),
	}
	if confInfo.ConsoleEnable {
		coreTeeArr = append(coreTeeArr, zapcore.NewCore(
			zapcore.NewConsoleEncoder(loggerConf.EncoderConfig),
			zapcore.AddSync(os.Stdout),
			loggerConf.Level,
		))
	}
	core := zapcore.NewTee(coreTeeArr...)

	options = append(options, zap.AddCaller(), zap.AddCallerSkip(0))
	Logger = zap.New(core, options...)
	return
}

func getWriteSyncer() (writeSyncer zapcore.WriteSyncer, err error) {
	logDir := path.Join(rootDir, confInfo.FilePath)
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}
	lumberjack := &lumberjack.Logger{
		Filename:   path.Join(logDir, confInfo.FileName),
		MaxSize:    confInfo.MaxSize,
		MaxBackups: confInfo.MaxBackups,
		MaxAge:     confInfo.MaxAge,
		Compress:   confInfo.IsCompress,
	}
	writeSyncer = zapcore.AddSync(lumberjack)
	return
}

func getConfig() (config zap.Config) {
	config = zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(getLevel())
	config.EncoderConfig = getEncodeConfig()
	return config
}

func getEncodeConfig() (config zapcore.EncoderConfig) {
	config = zap.NewProductionEncoderConfig()
	config.TimeKey = "time"
	config.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(time.Now().Format("2006-01-02 15:04:05.000"))
	}
	return config
}

// 配置文件的level转换为zapcore的level
func getLevel() (level zapcore.Level) {
	switch global.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	return
}
