package service

import (
	userRepo "gin_api_server_template/app/repository/user"
)

type UserService struct {
	UserInfoRepo userRepo.UserInfoRepo
}

func NewUserService(UserInfoRepo userRepo.UserInfoRepo) *UserService {
	return &UserService{
		UserInfoRepo: UserInfoRepo,
	}
}
