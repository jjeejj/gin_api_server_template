package response

const (
	ErrCode_Success    = 0
	ErrCode_InteralErr = -1
	ErrCode_ParamErr   = 100000 // 参数错误
	ErrCode_UnAuthErr  = 100001 // 未授权
)

var ErrMsgMap = map[int]string{
	ErrCode_Success:    "success",
	ErrCode_InteralErr: "internal error",
	ErrCode_ParamErr:   "param error",
	ErrCode_UnAuthErr:  "unauthorized",
}

// 额外的错误信息 ErrExtraBusinessMsg
const (
	PasswordNotAllowEmpty = "password not allow empty"  // 密码不允许为空
	AccountOrPasswordErr  = "account or password error" // 账号或密码错误
	VerifyCodeErr         = "verify code error"         // 验证码错误
	OldPasswordErr        = "old password error"        // 旧密码错误
)
