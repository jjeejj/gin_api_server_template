package cmd

import (
	"gin_api_server_template/app"
	"gin_api_server_template/app/route"
	"gin_api_server_template/internal/server"
)

func Run() {
	http := server.New()
	app.Start()
	http.GenRouter(route.New()) // 注册路由
	http.Run()
}
