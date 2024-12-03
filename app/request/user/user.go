package user

// UserLoginReq 登录
type UserLoginReq struct {
	// 登录账号
	Account string `json:"account" form:"account" binding:"required"`
	// 区号,如果是手机号 传 区号，如果是邮箱 传 邮箱服务商 gmail、qq、outlook 等
	AreaCode string `json:"area_code" form:"area_code" binding:"required"`
	// 登录密码，前端 必须进行 sha256 hash 后传过来 【小写】
	Password string `json:"password" form:"password"`
}

type GetUserInfoReq struct{}
