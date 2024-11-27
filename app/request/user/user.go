package user

// UserLoginReq 登录
type UserLoginReq struct {
	// 登录账号
	Account string `json:"account" form:"account" binding:"required"`
	// 区号,如果是手机号 传 区号，如果是邮箱 传 邮箱服务商 gmail、qq、outlook 等
	AreaCode string `json:"area_code" form:"area_code" binding:"required"`
	// 登录密码，前端 必须进行 sha256 hash 后传过来 【小写】
	// 1. 如果账号存在的时候登录，密码 或者验证码必须有一个
	// 2. 如果获取 验证码登录的时候，获取验证码返回的是新用户，需要设置密码，也是使用该字段
	Password string `json:"password" form:"password"`
	// 验证码
	Code string `json:"code" form:"code"`
}
