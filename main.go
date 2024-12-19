package main

import (
	"gin_api_server_template/cmd"
	"gin_api_server_template/internal/global"
	"gin_api_server_template/internal/logger"
)

func main() {
	cmd.Run()
	defer global.Rdb.Close()
	defer logger.Sync()
}
