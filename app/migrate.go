package app

import (
	"gin_api_server_template/app/model/user"

	"gorm.io/gorm"
)

func initModels(db *gorm.DB) {
	db.AutoMigrate(&user.UserInfo{})
}
