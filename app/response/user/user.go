package user

type ThirdUserLoginResp struct {
	// 用户登录成功的 token
	Token string `json:"token"`
	// 是否是新用户
	IsNew bool `json:"is_new"`
}

type UserLoginResp struct {
	// 用户登录成功的 token
	Token string `json:"token"`
	// 是否是新用户
	IsNew bool `json:"is_new"`
	// uid 用户id
	UserSysId string `json:"user_sys_id"`
}
