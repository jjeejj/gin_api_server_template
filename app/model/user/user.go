package user

import "gin_api_server_template/app/model"

// UserInfo 基础信息
type UserInfo struct {
	model.BaseModel
	SysId        string `json:"sys_id" gorm:"column:sys_id;uniqueIndex;comment:系统用户ID"`
	NickName     string `json:"nick_name" gorm:"column:nick_name;type:varchar(256);comment:昵称"`
	Avatar       string `json:"avatar" gorm:"column:avatar;type:varchar(256);comment:头像"`
	Sex          string `json:"sex" gorm:"column:sex;type:varchar(128);comment:性别 字典表:sex male:男;female:女; unknown:未知"`
	Birthday     string `json:"birthday" gorm:"column:birthday;type:varchar(128);comment:生日 年-月-日"`
	Introduction string `json:"introduction" gorm:"column:introduction;type:varchar(1024);comment:个人简介"`
	CountryCode  string `json:"country_code" gorm:"column:country_code;type:varchar(128);comment:国家编码"`
	LastLoginIp  string `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(128);comment:最近登录IP"`
	LastLoginAt  int64  `json:"last_login_at" gorm:"column:last_login_at;type:bigint(20);comment:最近登录时间"`
}
