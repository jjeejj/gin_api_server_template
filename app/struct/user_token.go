package _struct

import (
	"fmt"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/duke-git/lancet/v2/cryptor"
)

// UserToken 用户登录返回的 token
type UserToken struct {
	Token     string `json:"token"`
	UserSysId string `json:"user_sys_id"`
	ExpireAt  int64  `json:"expire_at"` // 过期时间
	CreateAt  int64  `json:"create_at"` // 创建时间
}

// SetExpireAt 设置过期时间
func (u *UserToken) SetExpireAt(expireAt int64) {
	u.ExpireAt = expireAt
}

func (u *UserToken) AddExpireTime(t int64) {
	u.ExpireAt = u.ExpireAt + t
}

// AesCbcEncrypt 加密 token
// 加密的没有 token, 生成成功后 再赋值
func (u *UserToken) AesCbcEncrypt() {
	userTokenBytes, _ := sonic.Marshal(u)
	resultByte := cryptor.AesCbcEncrypt(userTokenBytes, u.getAesKey())
	u.Token = cryptor.Base64StdEncode(string(resultByte))
}

// AesCbcDecrypt 解密 token
func (u *UserToken) AesCbcDecrypt() {
	userTokenBytes := cryptor.AesCbcDecrypt([]byte(cryptor.Base64StdDecode(u.Token)), u.getAesKey())
	sonic.Unmarshal(userTokenBytes, u)
}

// getAesKey 获取 加解密的 key
func (u *UserToken) getAesKey() []byte {
	return []byte(fmt.Sprintf("%s%s", u.UserSysId, strings.Repeat("0", 16)))[0:16]
}
