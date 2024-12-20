package logger

import (
	"context"
	"fmt"
	"gin_api_server_template/internal/config"

	"go.uber.org/zap"
)

type LogWrapper struct {
	logger *zap.Logger
}

var Log *LogWrapper = &LogWrapper{}

// 根据gin的context获取context，使log trace更加通用
// func GetTraceCtx(c *gin.Context) context.Context {
// 	return c.MustGet(consts.TraceCtx).(context.Context)
// }

func Init(confInfo config.LogConfig, rootDir string) (err error) {
	Log.logger, err = new(confInfo, rootDir)
	return
}

func Debugf(format string, msg ...any) {
	Log.logger.Debug(fmt.Sprintf(format, msg...))
}

func Warnf(format string, msg ...any) {
	Log.logger.Warn(fmt.Sprintf(format, msg...))
}

func Errorf(format string, msg ...any) {
	Log.logger.Error(fmt.Sprintf(format, msg...))
}

func Fatalf(format string, msg ...any) {
	Log.logger.Fatal(fmt.Sprintf(format, msg...))
}

func Infof(format string, msg ...any) {
	Log.logger.Info(fmt.Sprintf(format, msg...))
}
func DebugfCtx(ctx context.Context, format string, msg ...any) {
	trace := GetTraceCtx(ctx)
	Log.logger.Debug(fmt.Sprintf(format, msg...), trace.ToZapFields()...)
}

func WarnfCtx(ctx context.Context, format string, msg ...any) {
	trace := GetTraceCtx(ctx)
	Log.logger.Warn(fmt.Sprintf(format, msg...), trace.ToZapFields()...)
}

func ErrorfCtx(ctx context.Context, format string, msg ...any) {
	trace := GetTraceCtx(ctx)
	Log.logger.Error(fmt.Sprintf(format, msg...), trace.ToZapFields()...)
}

func FatalfCtx(ctx context.Context, format string, msg ...any) {
	trace := GetTraceCtx(ctx)
	Log.logger.Fatal(fmt.Sprintf(format, msg...), trace.ToZapFields()...)
}

func InfofCtx(ctx context.Context, format string, msg ...any) {
	trace := GetTraceCtx(ctx)
	Log.logger.Info(fmt.Sprintf(format, msg...), trace.ToZapFields()...)
}

func Sync() {
	Log.logger.Sync()
}
