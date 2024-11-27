package main

import (
	"gin_api_server_template/cmd"
	"gin_api_server_template/internal/global"
)

func main() {
	cmd.Run()
	defer global.Rdb.Close()
	defer global.Logger.Sync()
}
