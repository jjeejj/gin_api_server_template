package api

import "gin_api_server_template/app/api/admin"

var (
	AppApi  = &appApi{}
	UserApi = &userApi{}
	// 如果接口需要过多，需要单独的文件夹进行管理的时候，使用该方式就是初始化
	AdminApi = admin.NewAdminApi()
)
