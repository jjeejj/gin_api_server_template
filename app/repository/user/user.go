package user

import (
	"context"
	"gin_api_server_template/app/model/user"

	"gorm.io/gorm"
)

type UserInfoRepo interface {
	Create(ctx context.Context, user *user.UserInfo) error
	GetBySysId(ctx context.Context, sysId string) (*user.UserInfo, error) // 根据sysId获取用户信息
}

type DbUserInfoRepo struct {
	Db *gorm.DB
}

func NewDbUserInfoRepo(db *gorm.DB) UserInfoRepo {
	return &DbUserInfoRepo{
		Db: db,
	}
}

func (ui *DbUserInfoRepo) Create(ctx context.Context, user *user.UserInfo) error {
	return ui.Db.WithContext(ctx).Create(user).Error
}

func (ui *DbUserInfoRepo) CreateWithTrans(ctx context.Context, tx *gorm.DB, user *user.UserInfo) error {
	return tx.WithContext(ctx).Create(user).Error
}

func (ui *DbUserInfoRepo) GetBySysId(ctx context.Context, sysId string) (*user.UserInfo, error) {
	var userInfo *user.UserInfo
	err := ui.Db.WithContext(ctx).Model(&user.UserInfo{}).Where("sys_id = ?", sysId).First(&userInfo).Error
	return userInfo, err
}
