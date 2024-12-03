package user

type UserLoginResp struct {
	// 用户登录成功的 token
	Token string `json:"token"`
	// 是否是新用户
	IsNew bool `json:"is_new"`
	// uid 用户id
	UserSysId string `json:"user_sys_id"`
}

type GetUserInfoResp struct {
	// 昵称
	NickName string `json:"nick_name"`
	// 头像
	Avatar string `json:"avatar"`
	// 性别 male:男;female:女; unknown:未知 保密
	Sex string `json:"sex"`
	// 生日年-月-日
	Birthday string `json:"birthday"`
	// uid 用户id
	UserSysId string `json:"user_sys_id"`
	// 个人简介
	Introduction string `json:"introduction"`
}
