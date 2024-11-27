package user

import (
	"context"
	"gin_api_server_template/app/model/user"

	"gorm.io/gorm"
)

type UserInfoRepo interface {
	Create(ctx context.Context, user *user.UserInfo) error
	CreateWithTrans(ctx context.Context, tx *gorm.DB, user *user.UserInfo) error
	CheckSysIdIsExist(ctx context.Context, sysId string) (bool, error)
	GetBySysId(ctx context.Context, sysId string) (*user.UserInfo, error) // 根据sysId获取用户信息
	UpdateBySysId(ctx context.Context, sysId string, user *user.UserInfo) error
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

func (ui *DbUserInfoRepo) CheckSysIdIsExist(ctx context.Context, sysId string) (bool, error) {
	var count int64
	err := ui.Db.WithContext(ctx).Model(&user.UserInfo{}).Where("sys_id = ?", sysId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ui *DbUserInfoRepo) GetBySysId(ctx context.Context, sysId string) (*user.UserInfo, error) {
	var userInfo user.UserInfo
	err := ui.Db.WithContext(ctx).Model(&user.UserInfo{}).Where("sys_id = ?", sysId).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (ui *DbUserInfoRepo) UpdateBySysId(ctx context.Context, sysId string, userInfo *user.UserInfo) error {
	return ui.Db.WithContext(ctx).Model(&user.UserInfo{}).Where("sys_id = ?", sysId).Updates(userInfo).Error
}
