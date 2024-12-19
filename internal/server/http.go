package server

import (
	"context"
	_const "gin_api_server_template/internal/const"
	"gin_api_server_template/internal/global"
	"gin_api_server_template/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	_ "gin_api_server_template/internal/bootstrap"
)

type Http struct {
	engine *gin.Engine // http server
	port   string      // http server port
}

func New() *Http {
	if global.Config.App.Env == _const.Env_Prod {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Http{
		engine: gin.New(),
		port:   global.Config.App.Port,
	}
}

// GenRouter 注册路由
func (h *Http) GenRouter(r RouterGeneratorInterface) {
	r.AddRoute(h.engine)
}

func (h *Http) Run() {
	srv := &http.Server{
		Addr:    h.port,
		Handler: h.engine,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	// 监听服务退出信号
	h.ListenSignal(srv)
}

func (h *Http) ListenSignal(srv *http.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown ")
	}
}

func (h *Http) GetEngine() *gin.Engine {
	return h.engine
}
